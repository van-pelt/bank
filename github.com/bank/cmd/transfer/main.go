package main

import (
	"bank/pkg/bank/transfer"
	"flag"
	"fmt"
)

func main() {

	cmdFlagAmount := flag.Float64("amount", 5000, " a float")
	flag.Parse()
	if *cmdFlagAmount < 0 {
		fmt.Println("INVALID AMOUNT")
		return
	}

	total := transfer.Total(*cmdFlagAmount)
	fmt.Println("К зачислению", total, "дирам (", total/100, ") сомони")
}
