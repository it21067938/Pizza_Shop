package models

type Bill struct {
	BillID      int        `json:"billID"`
	TotalAmount float64    `json:"totalAmount"`
	Date        string     `json:"date"`
	Items       []BillItem `json:"items"`
}
