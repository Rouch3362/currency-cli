package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
)


func CallApi(currencyCode string) (map[string]APIResponse , *Error) {
	godotenv.Load()
	apiKey := os.Getenv("API_KEY")
	
	// call the api
	url  := fmt.Sprintf("http://api.navasan.tech/latest/?api_key=%s&item=%s",apiKey,currencyCode)
	res , err := http.Get(url)

	if err != nil {
		log.Fatal(err)
	}
	// close the body
	defer res.Body.Close()

	// it is map because of key's of currencies could be anything
	var result map[string]APIResponse

	
	err = json.NewDecoder(res.Body).Decode(&result)

	if err != nil {
		return nil , &Error{currencyCode , 404, time.Now().Format(time.RFC3339)}
	}
	

	return result , nil
}	


// iterates over multiple currencies
func IterateOverCurrencies(currencieCodes []string) []Response {
	// interface instance
	res := []Response{}
	for _,currencyCode := range currencieCodes {
		result, err := CallApi(currencyCode)
		
		if err != nil {
			res = append(res , err)
		} else {
			currencyRes := result[currencyCode]
			res = append(res,&currencyRes)
		}
	}

	return res
}
