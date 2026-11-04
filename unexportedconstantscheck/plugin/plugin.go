package plugin

import (
	"github.com/golangci/plugin-module-register/register"
	"golang.org/x/tools/go/analysis"

	"github.com/manuelarte/workshop-create-your-first-linter/unexportedconstantscheck/analyzer"
)

//nolint:gochecknoinits // init needed for plugin
func init() {
	register.Plugin("unexportedconstantscheck", New)
}

func New(_ any) (register.LinterPlugin, error) {
	return &unexportedConstantsCheckPlugin{}, nil
}

var _ register.LinterPlugin = new(unexportedConstantsCheckPlugin)

type unexportedConstantsCheckPlugin struct{}

func (u unexportedConstantsCheckPlugin) BuildAnalyzers() ([]*analysis.Analyzer, error) {
	return []*analysis.Analyzer{
		analyzer.New(),
	}, nil
}

func (u unexportedConstantsCheckPlugin) GetLoadMode() string {
	return register.LoadModeSyntax
}
