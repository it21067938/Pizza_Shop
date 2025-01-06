package controller

import (
	"backend/internal/db"
	"backend/internal/models"
	"encoding/json"
	"net/http"
)

// Get All Bills
func GetAllBills(w http.ResponseWriter, r *http.Request) {
	rows, err := db.DB.Query("SELECT BillID, TotalAmount, Date FROM bills")
	if err != nil{
		http.Error(w, "Failed to fetch items", http.StatusInternalServerError)
		return
	}

	defer rows.Close()

	var allBill []models.Bill
	for rows.Next(){
		var bill models.Bill
		if err := rows.Scan(&bill.BillID, &bill.TotalAmount, &bill.Date); err != nil {
			http.Error(w, "Failed to scan items", http.StatusInternalServerError)
			return
		}
		allBill = append(allBill, bill)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(allBill)
}
