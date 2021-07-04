package main

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/doc"
	"go/printer"
	"go/token"
	"log"
	"strings"

	"golang.org/x/tools/go/packages"
)

func display(pkg *doc.Package, fset *token.FileSet) {
	for _, cn := range pkg.Consts {
		fmt.Println(pkg.Name, "const", cn.Names, cn.Decl)
	}
	for _, vr := range pkg.Vars {
		fmt.Println(pkg.Name, "var", vr.Names, vr.Decl)
	}
	for _, tp := range pkg.Types {

		var kind string

		switch tp.Decl.Specs[0].(*ast.TypeSpec).Type.(type) {
		case *ast.InterfaceType:
			kind = "interface"
		default:
			log.Fatal("not implemented")
		}

		fmt.Println(pkg.Name, "type", tp.Name, kind)

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

		// interface methods
		for _, mt := range tp.Decl.Specs[0].(*ast.TypeSpec).Type.(*ast.InterfaceType).Methods.List {
			var params []string
			for _, param := range mt.Type.(*ast.FuncType).Params.List {
				var buf bytes.Buffer
				printer.Fprint(&buf, fset, param.Type)
				params = append(params, buf.String())
			}

			var results []string
			for _, result := range mt.Type.(*ast.FuncType).Results.List {
				var buf bytes.Buffer
				printer.Fprint(&buf, fset, result.Type)
				results = append(results, buf.String())
			}

			var returnType string
			switch len(results) {
			case 0:
				returnType = ""
			case 1:
				returnType = results[0]
			default:
				returnType = fmt.Sprintf("(%s)", strings.Join(results, ", "))
			}

			if len(mt.Names) > 1 {
				log.Fatal("interface methods has multiple names")
			}
			fullName := fmt.Sprintf("%s.%s", tp.Name, mt.Names[0])
			signature := fmt.Sprintf("%s(%s) %s", fullName, strings.Join(params, ", "), returnType)
			fmt.Println(pkg.Name, "func", signature)
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
		display(docpkg, pkg.Fset)
	}
}
