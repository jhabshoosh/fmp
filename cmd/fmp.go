package main

import (
	"fmt"
	"github.com/jhabshoosh/fmp/pkg/client"
)

func main() {
	ratios, err := client.FetchCashFlowStatements("AAPL")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(ratios)
}
