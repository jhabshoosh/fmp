package main

import (
	"fmt"
	"github.com/jhabshoo/fmp/client"
)

func main() {
	ratios, err := client.FetchFinancialStatements("AAPL")
	if (err != nil) {
		fmt.Println(err)
	}
	fmt.Println(ratios)
}