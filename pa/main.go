package main

import (
	"github.com/reusee/pa"
	"golang.org/x/tools/go/analysis/multichecker"
)

func main() {
	multichecker.Main(pa.Analyzers...)
}
