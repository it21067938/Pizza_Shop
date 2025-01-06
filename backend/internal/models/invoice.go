package models

import "time"

type Invoice struct {
	InvoiceID int         `json:"InvoiceID`
	Date      time.Time   `json:"Date`
	CartData  interface{} `json:CartData`
	Total     float64     `json:Total`
}
