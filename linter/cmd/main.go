package cmd

import (
	"fmt"
	"linter/internal"
)

func main() {
	fmt.Println("Starting Golang Linter...")
	internal.AnalyzeFile("testdata/sample.go") // test file for checking
}
