package main

import (
	"github.com/jhabshoo/fmp/pkg/client"
	"fmt"
)

func main() {
	ratios, err := client.FetchCashFlowStatements("AAPL")
	if (err != nil) {
		fmt.Println(err)
	}
	fmt.Println(ratios)
}