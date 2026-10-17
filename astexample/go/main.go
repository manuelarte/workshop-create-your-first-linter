package main

import (
	"go/ast"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() {
	singlechecker.Main(NewAnalyzer())
}

func NewAnalyzer() *analysis.Analyzer {
	return &analysis.Analyzer{
		Name:     "astexample",
		Doc:      "astexample to show the AST hierarchy",
		Run:      run,
		Requires: []*analysis.Analyzer{inspect.Analyzer},
	}
}

func run(pass *analysis.Pass) (any, error) {
	for _, file := range pass.Files {
		_ = ast.Print(pass.Fset, file)
	}
	return nil, nil
}
