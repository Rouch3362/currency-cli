package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(os.Args) < 2 {
		args = []string{"btc","eth","usdt"}
		fmt.Println("If you do not enter any currency code the defaults are: BTC,ETH,USDT")
	}

	
	responses := IterateOverCurrencies(args)

	for index, response := range responses {
		response.ShowMessage(args[index])
	}
}