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
		for _, cn := range docpkg.Consts {
			fmt.Println(pkg.Name, "const", cn.Names, cn.Decl)
		}
		for _, vr := range docpkg.Vars {
			fmt.Println(pkg.Name, "var", vr.Names, vr.Decl)
		}
		for _, tp := range docpkg.Types {
			fmt.Println(pkg.Name, "type", tp.Name)
		}
		for _, fn := range docpkg.Funcs {
			fmt.Println(pkg.Name, "func", fn.Name, fn.Decl)
		}

	}
}
