package models

type BillItem struct {
    BillItemID int     `json:"billItemID"`
    BillID     int     `json:"billID"`
    ItemID     int     `json:"itemID"`
    Quantity   int     `json:"quantity"`
    Price      float64 `json:"price"`
}
