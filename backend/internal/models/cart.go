package models

type Cart struct {
	ItemID   int     `json:"ItemID"`
    Quantity int     `json:"Quantity"`
    Item     Item    `json:"Item"`
}
