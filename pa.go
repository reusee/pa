package pa

import (
	"golang.org/x/tools/go/analysis"
	"honnef.co/go/tools/simple"
	"honnef.co/go/tools/staticcheck"
	"honnef.co/go/tools/unused"
)

var Analyzers = func() []*analysis.Analyzer {
	ret := append(XToolsAnalyzers[:0:0], XToolsAnalyzers...)
	for _, analyzer := range staticcheck.Analyzers {
		ret = append(ret, analyzer)
	}
	for _, analyzer := range simple.Analyzers {
		ret = append(ret, analyzer)
	}
	ret = append(ret, unused.Analyzer)
	return ret
}()
