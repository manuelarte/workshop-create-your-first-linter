package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"log/slog"
	"path/filepath"
)

func main() {
	logger := slog.Default()

	if err := run(); err != nil {
		logger.Error("Failed to run", "error", err)
	}
}

func run() error {
	path := filepath.Join("testdata", "main.go")
	fset := token.NewFileSet()
	file, err := parser.ParseFile(
		fset,
		path,
		nil,
		parser.ParseComments,
	)
	if err != nil {
		return fmt.Errorf("failed to parse file: %w", err)
	}

	if errPrinting := ast.Print(fset, file); errPrinting != nil {
		return fmt.Errorf("failed to print AST: %w", errPrinting)
	}

	return nil
}
