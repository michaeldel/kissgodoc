package main

import (
	"fmt"
	"go/doc"

	"golang.org/x/tools/go/packages"
)

func main() {
	cfg := &packages.Config{
		Mode: packages.LoadSyntax,
	}
	pkgs, err := packages.Load(cfg, "std")
	if err != nil {
		panic(err)
	}

	for _, pkg := range pkgs {
		fmt.Println(pkg)
		docpkg, err := doc.NewFromFiles(pkg.Fset, pkg.Syntax, pkg.Name)
		if err != nil {
			continue // TODO: is it right ?
		}
		for _, fn := range docpkg.Funcs {
			fmt.Println("  ", fn.Name, fn.Decl)
		}

	}
}
