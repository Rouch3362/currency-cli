package main

import (
	"fmt"
	"strconv"

	"github.com/dustin/go-humanize"
	"github.com/fatih/color"
)

type Response interface {
	ShowMessage(string)
}

type APIResponse struct {
	Price  		string  `json:"value"`
	Change 		float32 `json:"change"`
	Date   		string  `json:"date"`
}

func (a *APIResponse) ShowMessage(currencyCode string) {
	// convert price string to int64
	priceInt , err := strconv.ParseInt(a.Price,10,64)

	var PriceInHumanReadableFormat string
	// make price in human readable format
	if err != nil {
		PriceInHumanReadableFormat = a.Price
		err = nil
	} else {
		PriceInHumanReadableFormat = humanize.Comma(priceInt)
	}

	priceChangeHumaneReadable := humanize.Comma(int64(a.Change))

	// format a message
	message := fmt.Sprintf("Price Of %s at %s - %sT with %sT of Changes." ,currencyCode, a.Date , PriceInHumanReadableFormat , priceChangeHumaneReadable)
	
	// if price goes up
	if a.Change > 0 {
		color.Blue("▲ %s" , message)
	} else if a.Change < 0 { // if price goes down
		color.Red("▼ %s" , message)
	} else { // if price stays
		color.White("- %s" , message)
	}
	
}

type Error struct {
	Currency string
	Code     int32
	Date     string
}

func (e *Error) ShowMessage(currencyCode string) {
	color.Red("- Price of %s you entered not found %d." , e.Currency, e.Code)
}