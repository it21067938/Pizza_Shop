package models

type Item struct {
    ItemID   int     `json:"ItemID"`
    Name     string  `json:"Name"`
    Price    float64 `json:"Price"`
    Category string  `json:"Category"`
}
