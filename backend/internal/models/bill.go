package models

type Bill struct {
	BillID      int     `json:"BillID"`
	TotalAmount float64 `json:"TotalAmount"`
	CartItems   []Cart  `json:"CartItems"`
	Date        string  `json:"Date`
}
