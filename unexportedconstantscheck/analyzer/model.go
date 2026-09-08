package analyzer

import (
	"go/ast"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
)

func New() *analysis.Analyzer {
	return &analysis.Analyzer{
		Name:     "unexportedconstantscheck",
		Doc:      "unexportedconstantscheck checks if unexported constants starts with _",
		Run:      run,
		Requires: []*analysis.Analyzer{inspect.Analyzer},
	}
}

func run(pass *analysis.Pass) (any, error) {

	for _, file := range pass.Files {
		for _, decl := range file.Decls {
			// Exercise
			_ = ast.Print(pass.Fset, decl)
		}
	}

	return nil, nil
}
