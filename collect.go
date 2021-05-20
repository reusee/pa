// +build ignore

// collect all *analysis.Analyzer in golang.org/x/tools/go/analysis/passes/...

package main

import (
	"fmt"
	"go/format"
	"go/types"
	"io/ioutil"
	"sort"
	"strings"

	"golang.org/x/tools/go/packages"
)

var (
	pt = fmt.Printf
)

func main() {

	pkgs, err := packages.Load(
		&packages.Config{
			Mode: packages.NeedTypesInfo |
				packages.NeedFiles |
				packages.NeedSyntax |
				packages.NeedTypes |
				packages.NeedName,
		},
		"golang.org/x/tools/go/analysis/...",
	)
	if err != nil {
		panic(err)
	}
	if packages.PrintErrors(pkgs) > 0 {
		return
	}

	var analyzerType types.Type
	packages.Visit(pkgs, nil, func(pkg *packages.Package) {
		if pkg.PkgPath != "golang.org/x/tools/go/analysis" {
			return
		}
		analyzerType = types.NewPointer(
			pkg.Types.Scope().Lookup("Analyzer").Type(),
		)
	})
	if analyzerType == nil {
		panic("analysis.Analyzer not found")
	}

	analyzers := make(map[string]map[string]bool)
	packages.Visit(pkgs, nil, func(pkg *packages.Package) {
		if map[string]bool{
			// too many false positives
			"golang.org/x/tools/go/analysis/passes/shadow": true,
		}[pkg.PkgPath] {
			return
		}
		globalScope := pkg.Types.Scope()
		for _, name := range globalScope.Names() {
			obj := globalScope.Lookup(name)
			if !types.Identical(obj.Type(), analyzerType) {
				return
			}
			m, ok := analyzers[pkg.PkgPath]
			if !ok {
				m = make(map[string]bool)
				analyzers[pkg.PkgPath] = m
			}
			m[fmt.Sprintf("%s.%s", pkg.Name, name)] = true
			pt("%v\n", obj)
		}
	})

	src := `package pa

import (
  "golang.org/x/tools/go/analysis"
  ` + func() (ret string) {
		var paths []string
		for path := range analyzers {
			paths = append(paths, fmt.Sprintf(`"%s"`, path))
		}
		sort.Strings(paths)
		return strings.Join(paths, "\n")
	}() + `
)

var XToolsAnalyzers = []*analysis.Analyzer{
  ` + func() (ret string) {
		var names []string
		for _, m := range analyzers {
			for name := range m {
				names = append(names, name+",")
			}
		}
		sort.Strings(names)
		return strings.Join(names, "\n")
	}() + `
}
`

	bs, err := format.Source([]byte(src))
	if err != nil {
		panic(err)
	}
	if err := ioutil.WriteFile("analyzers.go", bs, 0644); err != nil {
		panic(err)
	}

}
