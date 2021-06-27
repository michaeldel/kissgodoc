package main

import (
	"fmt"
	"go/ast"
	"go/doc"
	"reflect"

	"golang.org/x/tools/go/packages"
)

func display(pkg *doc.Package) {
	for _, cn := range pkg.Consts {
		fmt.Println(pkg.Name, "const", cn.Names, cn.Decl)
	}
	for _, vr := range pkg.Vars {
		fmt.Println(pkg.Name, "var", vr.Names, vr.Decl)
	}
	for _, tp := range pkg.Types {
		fmt.Println(pkg.Name, "type", tp.Name)
		fmt.Println(
			pkg.Name, "type", reflect.TypeOf(tp.Decl.Specs[0].(*ast.TypeSpec).Type),
		)

		for _, cn := range tp.Consts {
			fmt.Println(pkg.Name, "const", tp.Name, cn.Names, cn.Decl)
		}
		for _, vr := range tp.Vars {
			fmt.Println(pkg.Name, "var", tp.Name, vr.Names, vr.Decl)
		}
		for _, fn := range tp.Funcs {
			fmt.Println(pkg.Name, "func", tp.Name, fn.Name, fn.Decl)
		}
		for _, mt := range tp.Methods {
			fmt.Println(pkg.Name, "method", tp.Name, mt.Name, mt.Decl)
		}
	}
	for _, fn := range pkg.Funcs {
		fmt.Print(pkg.Name, " func ", fn.Name, "(")
		for i, field := range fn.Decl.Type.Params.List {
			if i != 0 {
				fmt.Print(", ")
			}
			// TODO: multiple names ?
			fmt.Print(field.Names[0], field.Type)
		}
		fmt.Print(")")

		if fn.Decl.Type.Results == nil {
			break
		}

		for i, field := range fn.Decl.Type.Results.List {
			if i == 0 {
				fmt.Print(" ")
			} else {
				fmt.Print(", ")
			}
			// TODO: named
			// TODO: multiple values parentheses
			// TODO: multiple names ?
			fmt.Print(field.Type)
		}
	}
}

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
		display(docpkg)
	}
}
