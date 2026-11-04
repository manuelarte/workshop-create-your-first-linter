package analyzer

import (
	"fmt"
	"go/ast"
	"go/token"
	"strings"

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
			// We are only interested in the GenDecl nodes that are constants
			constDecl, isGenDecl := decl.(*ast.GenDecl)
			if !isGenDecl || constDecl.Tok != token.CONST {
				continue
			}
			// Iterate over the ValueSpec nodes within the GenDecl
			for _, spec := range constDecl.Specs {
				valueSpec, ok := spec.(*ast.ValueSpec)
				if !ok {
					continue
				}
				// logic to check if the constant is unexported and starts with an underscore
				for _, name := range valueSpec.Names {
					if name.IsExported() {
						continue
					}
					if strings.HasPrefix(name.Name, "err") {
						continue
					}
					if name.Name[0] == '_' {
						continue
					}
					pass.Report(unexportedConstantWithoutUnderscore(name))
				}
			}
		}
	}

	return nil, nil
}

func unexportedConstantWithoutUnderscore(name *ast.Ident) analysis.Diagnostic {
	return analysis.Diagnostic{
		Pos:     name.Pos(),
		End:     name.End(),
		Message: fmt.Sprintf(`unexported constant %q should start with "_"`, name.Name),
	}
}
