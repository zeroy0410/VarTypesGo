package main

import (
	"fmt"
	"github.com/zeroy0410/VarTypesGo/typeAnalyzer"
)

func main() {
	//runner := typeAnalyzer.NewRunner("C:\\Users\\zeroy\\Documents\\Code\\VarTypesGo\\case1\\...")
	//runner.Dir = "C:\\Users\\zeroy\\Documents\\Code\\VarTypesGo\\examples\\case1"
	//runner.ModuleName = "github.com/prometheus/prometheus"
	runner := typeAnalyzer.NewRunner("C:/Users/zeroy/Documents/Code/VarTypesGo/examples/case1/...")
	runner.Dir = "C:/Users/zeroy/Documents/Code/VarTypesGo/examples/case1"
	runner.ExportToSSA = true
	runner.AnalyzerName = "vta"
	err := runner.Run()
	if err != nil {
		fmt.Println(err)
	}
}
