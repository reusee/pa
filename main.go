package main

import (
	"golang.org/x/tools/go/analysis/multichecker"
	"honnef.co/go/tools/simple"
	"honnef.co/go/tools/staticcheck"
	"honnef.co/go/tools/unused"
)

func main() {
	for _, analyzer := range staticcheck.Analyzers {
		Analyzers = append(Analyzers, analyzer)
	}
	for _, analyzer := range simple.Analyzers {
		Analyzers = append(Analyzers, analyzer)
	}
	Analyzers = append(Analyzers, unused.Analyzer)
	multichecker.Main(Analyzers...)
}
