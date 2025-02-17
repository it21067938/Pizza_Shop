package controller

import (
	"backend/internal/db"
	"backend/internal/models"
	"encoding/json"
	//"log"
	"net/http"
)

func CreateBill(w http.ResponseWriter, r *http.Request) {
	var request struct {
		TotalAmount float64           `json:"totalAmount"`
		Items       []models.BillItem `json:"items"`
	}

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	tx, err := db.DB.Begin()
	if err != nil {
		http.Error(w, "Failed to start transaction", http.StatusInternalServerError)
		return
	}

	result, err := tx.Exec("INSERT INTO bills (TotalAmount, Date) VALUES (?, NOW())", request.TotalAmount)
	if err != nil {
		tx.Rollback()
		http.Error(w, "Failed to create bill", http.StatusInternalServerError)
		return
	}

	billID, _ := result.LastInsertId()

	for _, item := range request.Items {
		_, err := tx.Exec("INSERT INTO bill_items (BillID, ItemID, Quantity, Price) VALUES (?, ?, ?, ?)",
			billID, item.ItemID, item.Quantity, item.Price)
		if err != nil {
			//log.Printf("Failed to insert : %v", err)
			tx.Rollback()
			http.Error(w, "Failed to insert bill items", http.StatusInternalServerError)
			return
		}
	}

	if err := tx.Commit(); err != nil {
		http.Error(w, "Failed to commit transaction", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "Bill created successfully",
		"billID":  billID,
	})
}

// GetAllBills
func GetAllBills(w http.ResponseWriter, r *http.Request) {
	rows, err := db.DB.Query("SELECT BillID, TotalAmount, Date FROM bills")
	if err != nil {
		http.Error(w, "Failed to fetch bills", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var bills []models.Bill

	for rows.Next() {
		var bill models.Bill
		if err := rows.Scan(&bill.BillID, &bill.TotalAmount, &bill.Date); err != nil {
			http.Error(w, "Failed to read bill data", http.StatusInternalServerError)
			return
		}

		itemRows, err := db.DB.Query("SELECT ItemID, Quantity, Price FROM bill_items WHERE BillID = ?", bill.BillID)
		if err != nil {
			http.Error(w, "Failed to fetch bill items", http.StatusInternalServerError)
			return
		}

		var items []models.BillItem
		for itemRows.Next() {
			var item models.BillItem
			if err := itemRows.Scan(&item.ItemID, &item.Quantity, &item.Price); err != nil {
				http.Error(w, "Failed to read bill item data", http.StatusInternalServerError)
				return
			}
			items = append(items, item)
		}
		itemRows.Close()

		bill.Items = items
		bills = append(bills, bill)
	}

	if err := json.NewEncoder(w).Encode(bills); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}
