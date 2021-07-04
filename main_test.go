package main

import (
	"go/ast"
	"go/doc"
	"go/parser"
	"go/token"
)

func ExampleEmpty() {
	fset := token.NewFileSet()
	docpkg, err := doc.NewFromFiles(fset, []*ast.File{}, "foo")
	if err != nil {
		panic(err)
	}

	display(docpkg, fset)
	// Output:
}

func ExampleTrivialFunction() {
	src := `
package trivial

func F() {}
`
	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, "trivial.go", src, parser.ParseComments)
	if err != nil {
		panic(err)
	}

	docpkg, err := doc.NewFromFiles(fset, []*ast.File{file}, "trivial")
	if err != nil {
		panic(err)
	}

	display(docpkg, fset)
	// Output:
	// trivial func F()
}

func ExampleFibonacci() {
	src := `
package fibonacci

func Fib(n int32) int32 {
	if n == 1 || n == 2 {
		return 1
	}
	return Fib(n - 1) + fib(n - 2)
}
`
	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, "fib.go", src, parser.ParseComments)
	if err != nil {
		panic(err)
	}

	docpkg, err := doc.NewFromFiles(fset, []*ast.File{file}, "fib")
	if err != nil {
		panic(err)
	}

	display(docpkg, fset)
	// Output:
	// fibonacci func Fib(n int32) int32
}

func ExampleInterface() {
	src := `
package fuzzer

type Fuzzer interface {
	Fuzz(int, []int, map[string]int, struct{}, interface{}) interface{}
}
`
	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, "fuzzer.go", src, parser.ParseComments)
	if err != nil {
		panic(err)
	}

	docpkg, err := doc.NewFromFiles(fset, []*ast.File{file}, "fib")
	if err != nil {
		panic(err)
	}

	display(docpkg, fset)
	// Output:
	// fuzzer type Fuzzer interface
	// fuzzer func Fuzzer.Fuzz(int, []int, map[string]int, struct{}, interface{}) interface{}
}
