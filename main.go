package main

import (
	"fmt"
	"os"
)

func main() {
	args := []string{"btc","eth","usdt"}
	if len(os.Args) > 1 {
		args = os.Args[1:]
		fmt.Println("If you do not enter any currency code the defaults are: BTC,ETH,USDT")
	}

	
	responses := IterateOverCurrencies(args)

	for index, response := range responses {
		response.ShowMessage(os.Args[index+1])
	}
}