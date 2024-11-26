package main

import (
	"fmt"
	"github.com/zeroy0410/VarTypesGo/typeAnalyzer"
)

func main() {
	runner := typeAnalyzer.NewRunner("C:\\Users\\zeroy\\Documents\\Code\\testdata_for_VarTypesGo\\moby\\...")
	runner.Dir = "C:\\Users\\zeroy\\Documents\\Code\\testdata_for_VarTypesGo\\moby"
	runner.ModuleName = "github.com/docker/docker"
	//runner := typeAnalyzer.NewRunner("C:/Users/zeroy/Documents/Code/VarTypesGo/examples/case2/...")
	//runner.Dir = "C:/Users/zeroy/Documents/Code/VarTypesGo/examples/case2"
	runner.ExportToSSA = true
	runner.AnalyzerName = "kcfa"
	err := runner.Run()
	if err != nil {
		fmt.Println(err)
	}
}
