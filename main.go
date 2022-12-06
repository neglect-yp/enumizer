package main

import (
	"bytes"
	"context"
	"flag"
	"go/ast"
	"go/format"
	"go/token"
	"go/types"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"text/template"

	"github.com/iancoleman/strcase"
	"golang.org/x/tools/go/packages"
)

var (
	path   = flag.String("path", "./internal/api/domain/model", "path to model package")
	output = flag.String("output", "enumizer.gen.go", "filename to output")
)

type EnumPackages map[string]EnumPackage
type EnumPackage struct {
	Path  string
	Enums Enums
}

type Enums map[string]Enum

func (e *Enums) Put(name string, variants ...string) {
	if e == nil {
		*e = make(Enums, 0)
	}

	v, ok := (*e)[name]
	if !ok {
		(*e)[name] = Enum{Name: name, Variants: variants}
	} else {
		v.Variants = append(v.Variants, variants...)
		(*e)[name] = v
	}
}

func (e Enums) SortedEnums() []Enum {
	enums := make([]Enum, 0)
	for _, v := range e {
		enums = append(enums, v)
	}

	sort.SliceStable(enums, func(i, j int) bool {
		return enums[i].Name < enums[j].Name
	})

	return enums
}

type Enum struct {
	Name     string
	Variants []string
}

func FindEnumPackages(path string) (EnumPackages, error) {
	pkgs, err := packages.Load(&packages.Config{
		Mode: packages.NeedName |
			packages.NeedFiles |
			packages.NeedCompiledGoFiles |
			packages.NeedImports |
			packages.NeedTypes |
			packages.NeedTypesSizes |
			packages.NeedSyntax |
			packages.NeedTypesInfo,
	}, path)
	if err != nil {
		return nil, err
	}

	enumPackages := make(EnumPackages, 0)
	for _, pkg := range pkgs {
		enums := make(Enums, 0)
		for _, file := range pkg.Syntax {
			ast.Inspect(file, func(node ast.Node) bool {
				genDecl, ok := node.(*ast.GenDecl)
				if !ok {
					return true
				}

				if genDecl.Tok != token.CONST {
					return true
				}

				if genDecl.Doc == nil || !existsMarker(genDecl.Doc.List) {
					return true
				}

				var typeName string
				variants := make([]string, 0)
				for _, spec := range genDecl.Specs {
					spec, ok := spec.(*ast.ValueSpec)
					if !ok {
						continue
					}

					for _, name := range spec.Names {
						t := types.TypeString(pkg.TypesInfo.ObjectOf(name).Type(), types.RelativeTo(pkg.Types))
						if typeName == "" {
							typeName = t
						}
						if t != typeName {
							logWithPos(pkg.Fset, genDecl.Pos(), "[Error] target const decl must include only one type. ignored")
							return false
						}
						variants = append(variants, name.String())
					}
				}

				enums.Put(typeName, variants...)
				return false
			})
		}
		enumPackages[pkg.Name] = EnumPackage{
			Path:  pkg.PkgPath,
			Enums: enums,
		}
	}

	return enumPackages, nil
}

const marker = "enumizer:generate"

func existsMarker(comments []*ast.Comment) bool {
	for _, comment := range comments {
		commentBody := strings.TrimSpace(strings.TrimPrefix(comment.Text, "//"))
		if commentBody == marker {
			return true
		}
	}

	return false
}

func logWithPos(fset *token.FileSet, pos token.Pos, message string) {
	p := fset.Position(pos)
	log.Printf("%s:%d: %s", p.Filename, p.Line, message)
}

func GenerateEnumHelpers(packageName string, enums Enums) ([]byte, error) {
	templateArgs := struct {
		PackageName string
		Enums       Enums
	}{
		PackageName: packageName,
		Enums:       enums,
	}

	tpl := template.Must(template.New("").Funcs(map[string]interface{}{
		"lowerCamel": strcase.ToLowerCamel,
	}).Parse(`
// Code generated by enumizer; DO NOT EDIT.
package {{ .PackageName }}

import "fmt"

{{ range $i, $enum := .Enums.SortedEnums -}}
var {{ lowerCamel $enum.Name }}Set = map[{{ $enum.Name }}]struct{}{
{{ range $j, $variant := $enum.Variants -}}
{{ $variant }}: {},
{{ end -}}
}

func {{ $enum.Name }}List() []{{ $enum.Name }} {
	ret := make([]{{ $enum.Name }}, 0, len({{ lowerCamel $enum.Name }}Set))
	for v := range {{ lowerCamel $enum.Name }}Set {
		ret = append(ret, v)
	}
	return ret
}

func (m {{ $enum.Name }}) IsValid() bool {
	_, ok := {{ lowerCamel $enum.Name }}Set[m]
	return ok
}

func (m {{ $enum.Name }}) Validate() error {
	if !m.IsValid() {
		return fmt.Errorf("{{ $enum.Name }}(%v) is invalid", m)
	}
	return nil
}
{{ end -}}
`))

	buf := new(bytes.Buffer)
	if err := tpl.Execute(buf, templateArgs); err != nil {
		return nil, err
	}

	return format.Source(buf.Bytes())
}

func run(ctx context.Context) error {
	enumPackages, err := FindEnumPackages(*path)
	if err != nil {
		return err
	}

	for packageName, enumPackage := range enumPackages {
		src, err := GenerateEnumHelpers(packageName, enumPackage.Enums)
		if err != nil {
			return err
		}

		if err := os.WriteFile(filepath.Join(enumPackage.Path, *output), src, 0644); err != nil {
			return err
		}
	}

	return nil
}

func main() {
	flag.Parse()

	if err := run(context.Background()); err != nil {
		log.Fatalf("%+v", err)
	}
}
