package internal

import (
	"fmt"
	"go/ast"
)

func checkFunctionNames(fn *ast.FuncDecl) {
	// for example checking name of functions
	if fn.Name.IsExported() && fn.Name.Name[0] < 'A' || fn.Name.Name[0] > 'Z' {
		fmt.Printf("Warning: Function name '%s' should start with an uppercase letter\n", fn.Name.Name)
	}
}
