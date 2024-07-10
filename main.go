package main

import (
	"os"
)

func main() {

	responses := IterateOverCurrencies(os.Args[1:])
	
	for index, response := range responses {
		response.ShowMessage(os.Args[index+1])
	}
}