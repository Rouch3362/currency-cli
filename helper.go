package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)


func CallApi(currencyCode string) *map[string]APIResponse {
	godotenv.Load()
	apiKey := os.Getenv("API_KEY")
	
	url  := fmt.Sprintf("http://api.navasan.tech/latest/?api_key=%s&item=%s",apiKey,currencyCode)
	res , err := http.Get(url)

	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()

	var result map[string]APIResponse
	
	err = json.NewDecoder(res.Body).Decode(&result)

	if err != nil {
		log.Fatal(err)
	}

	return &result
}	