package main

import (
	"golang.org/x/tools/go/analysis/multichecker"
	"honnef.co/go/tools/simple"
	"honnef.co/go/tools/staticcheck"
)

func main() {
	for _, analyzer := range staticcheck.Analyzers {
		Analyzers = append(Analyzers, analyzer)
	}
	for _, analyzer := range simple.Analyzers {
		Analyzers = append(Analyzers, analyzer)
	}
	multichecker.Main(Analyzers...)
}
