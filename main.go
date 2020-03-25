package main

import (
	"golang.org/x/tools/go/analysis/multichecker"
	"honnef.co/go/tools/staticcheck"
)

func main() {
	for _, analyzer := range staticcheck.Analyzers {
		Analyzers = append(Analyzers, analyzer)
	}
	multichecker.Main(Analyzers...)
}
