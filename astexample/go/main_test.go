package main

import (
	"testing"

	"golang.org/x/tools/go/analysis/analysistest"
)

func TestAnalyzer(t *testing.T) {
	a := NewAnalyzer()

	analysistest.Run(t, analysistest.TestData(), a, "simple")
}
