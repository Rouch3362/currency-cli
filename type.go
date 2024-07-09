package main



type APIResponse struct {
	Price 	string	`json:"value"`
	Change 	int32 	`json:"change"`
	Date	string	`json:"date"`
}