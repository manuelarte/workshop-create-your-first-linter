package main

import (
	"golang.org/x/tools/go/analysis/singlechecker"

	"github.com/manuelarte/workshop-create-your-first-linter/unexportedconstantscheck/analyzer"
)

func main() {
	singlechecker.Main(analyzer.New())
}
