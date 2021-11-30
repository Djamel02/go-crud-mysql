package models

import "time"

// Employee detail
type Employee struct {
	ID      	int64  `json:"id"`
	Name    	string `json:"name"`
	Phone   	string `json:"phone"`
	Picture 	string `json:"picture"`
	Job			string	`json:"job"`
	Country 	string `json:"country"`
	City		string	`json:"city"`
	Postalcode 	int64 `json:"postalcode"`
	CreatedAt 	*time.Time `json:"created_at"`
}
