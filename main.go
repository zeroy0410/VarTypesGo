package main

import (
	"fmt"
	"github.com/zeroy0410/VarTypesGo/typeAnalyzer"
)

func main() {
	runner := typeAnalyzer.NewRunner("/Users/zeroy/Documents/Code/VarTypesGo/examples/case1/...")
	runner.Dir = "/Users/zeroy/Documents/Code/VarTypesGo/examples/case1/"
	runner.ExportToSSA = true
	runner.AnalyzerName = "vtafs"
	runner.K = 2
	err := runner.Run()
	if err != nil {
		fmt.Println(err)
	}
}
