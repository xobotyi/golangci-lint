package exhaustruct

import (
	"dev.gaijin.team/go/exhaustruct/v4/analyzer"
	"golang.org/x/tools/go/analysis"

	"github.com/golangci/golangci-lint/pkg/config"
	"github.com/golangci/golangci-lint/pkg/goanalysis"
	"github.com/golangci/golangci-lint/pkg/golinters/internal"
)

func New(settings *config.ExhaustructSettings) *goanalysis.Linter {
	cfg := analyzer.Config{}
	if settings != nil {
		cfg.IncludeRx = settings.Include
		cfg.ExcludeRx = settings.Exclude
		cfg.AllowEmpty = settings.AllowEmpty
		cfg.AllowEmptyRx = settings.AllowEmptyRx
		cfg.AllowEmptyReturns = settings.AllowEmptyReturns
		cfg.AllowEmptyDeclarations = settings.AllowEmptyDeclarations
	}

	a, err := analyzer.NewAnalyzer(cfg)
	if err != nil {
		internal.LinterLogger.Fatalf("exhaustruct configuration: %v", err)
	}

	return goanalysis.NewLinter(
		a.Name,
		a.Doc,
		[]*analysis.Analyzer{a},
		nil,
	).WithLoadMode(goanalysis.LoadModeTypesInfo)
}
