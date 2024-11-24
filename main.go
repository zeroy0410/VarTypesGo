package main

import (
	"fmt"
	"github.com/zeroy0410/VarTypesGo/typeAnalyzer"
	"time"
)

func main() {
	startTime := time.Now()
	runner := typeAnalyzer.NewRunner("C:/Users/zeroy/Documents/Code/go-tools/...")
	runner.Dir = "C:/Users/zeroy/Documents/Code/go-tools/"
	runner.ExportToSSA = true
	runner.AnalyzerName = "vtafs"
	err := runner.Run()
	endTime := time.Now()
	executionTime := endTime.Sub(startTime)
	fmt.Printf("Time： %s\n", executionTime)
	if err != nil {
		fmt.Println(err)
	}
}
