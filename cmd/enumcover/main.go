package main

import (
	"github.com/neglect-yp/enumizer/analyzer/enumcover"
	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() {
	singlechecker.Main(enumcover.Analyzer)
}
