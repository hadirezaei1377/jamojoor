package internal

import (
	"fmt"
	"go/parser"
	"go/token"
)

func AnalyzeFile(filename string) {
	fset := token.NewFileSet()

	node, err := parser.ParseFile(fset, filename, nil, parser.AllErrors)
	if err != nil {
		fmt.Printf("Error parsing file: %v\n", err)
		return
	}

	for _, decl := range node.Decls {
		checkNamingConventions(decl)
	}
}

func checkNamingConventions(decl interface{}) {
	// related logic
	fmt.Println("Checking naming conventions...")
}
