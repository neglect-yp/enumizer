package enumcover

import (
	"go/ast"
	"go/token"
	"go/types"
	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
	"sort"
	"strings"
)

var Analyzer = &analysis.Analyzer{
	Name:      "enumcover",
	Doc:       "check enum coverage",
	Requires:  []*analysis.Analyzer{inspect.Analyzer},
	Run:       run,
	FactTypes: []analysis.Fact{new(isEnum)},
}

type isEnum struct {
	variants []string
}

func (f *isEnum) AFact() {}

func (f *isEnum) String() string {
	return strings.Join(f.variants, ", ")
}

func (f *isEnum) variantsSet() map[string]struct{} {
	set := make(map[string]struct{})
	for _, variant := range f.variants {
		set[variant] = struct{}{}
	}
	return set
}

func run(pass *analysis.Pass) (interface{}, error) {
	findEnumTypes(pass)
	checkSwitchCoverage(pass)
	return nil, nil
}

func findEnumTypes(pass *analysis.Pass) {
	inspect := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)
	nodeFilter := []ast.Node{
		(*ast.GenDecl)(nil),
	}

	inspect.Preorder(nodeFilter, func(n ast.Node) {
		genDecl, ok := n.(*ast.GenDecl)
		if !ok {
			return
		}

		if genDecl.Tok != token.CONST {
			return
		}

		if genDecl.Doc == nil || !existsMarker(genDecl.Doc.List) {
			return
		}

		var typeName string
		var typ types.Type
		variants := make([]string, 0)
		for _, spec := range genDecl.Specs {
			spec, ok := spec.(*ast.ValueSpec)
			if !ok {
				continue
			}

			for _, name := range spec.Names {
				typ = pass.TypesInfo.ObjectOf(name).Type()
				t := types.TypeString(typ, types.RelativeTo(pass.Pkg))
				if typeName == "" {
					typeName = t
				}
				if t != typeName {
					pass.Reportf(genDecl.Pos(), "target const decl must include only one type")
					return
				}
				variants = append(variants, name.String())
			}
		}

		named, ok := typ.(*types.Named)
		if !ok {
			pass.Reportf(genDecl.Pos(), "includes not named type") // debug
			return
		}

		var fact isEnum
		if pass.ImportObjectFact(named.Obj(), &fact) {
			fact.variants = append(fact.variants, variants...)
		} else {
			fact = isEnum{variants: variants}
		}
		pass.ExportObjectFact(named.Obj(), &fact)
	})
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

func checkSwitchCoverage(pass *analysis.Pass) {
	inspect := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)
	nodeFilter := []ast.Node{
		(*ast.SwitchStmt)(nil),
	}

	inspect.Preorder(nodeFilter, func(n ast.Node) {
		switchStmt, ok := n.(*ast.SwitchStmt)
		if !ok {
			return
		}

		named, ok := pass.TypesInfo.TypeOf(switchStmt.Tag).(*types.Named)
		if !ok {
			return
		}

		var fact isEnum
		if !pass.ImportObjectFact(named.Obj(), &fact) {
			return
		}

		variantsSet := fact.variantsSet()
		for _, stmt := range switchStmt.Body.List {
			caseClause, ok := stmt.(*ast.CaseClause)
			if !ok {
				continue
			}

			names := make([]string, 0)
			for _, expr := range caseClause.List {
				switch expr := expr.(type) {
				case *ast.Ident:
					names = append(names, expr.Name)
				case *ast.SelectorExpr:
					names = append(names, expr.Sel.Name)
				}
			}

			for _, name := range names {
				delete(variantsSet, name)
			}
		}

		if len(variantsSet) > 0 {
			keys := make([]string, 0)
			for key := range variantsSet {
				keys = append(keys, key)
			}
			sort.Strings(keys)

			pass.Reportf(switchStmt.Pos(), "this switch statement doesn't cover enum variants. missing cases: %v", strings.Join(keys, ", "))
		}
	})
}
