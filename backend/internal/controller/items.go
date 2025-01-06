package controller

import (
	"backend/internal/db"
	"backend/internal/models"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// add item
func AddItem(w http.ResponseWriter, r *http.Request) {
	var item models.Item
	if err := json.NewDecoder(r.Body).Decode(&item); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	result, err := db.DB.Exec("INSERT INTO items (Name, Price, Category) VALUES (?, ?, ?)", item.Name, item.Price, item.Category)
	if err != nil {
		http.Error(w, "Failed to create item", http.StatusInternalServerError)
		return
	}

	id, _ := result.LastInsertId()
	item.ItemID = int(id)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(item)

}

// Get all items
func GetItems(w http.ResponseWriter, r *http.Request) {
	rows, err := db.DB.Query("SELECT ItemID, Name, Price, Category FROM items")
	if err != nil {
		http.Error(w, "Failed to fetch items", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var items []models.Item
	for rows.Next() {
		var item models.Item
		if err := rows.Scan(&item.ItemID, &item.Name, &item.Price, &item.Category); err != nil {
			http.Error(w, "Failed to scan items", http.StatusInternalServerError)
			return
		}
		items = append(items, item)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(items)
}

// Delete a item
func DeleteItem(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	ItemID, _ := strconv.Atoi(params["ItemID"])

	_, err := db.DB.Exec("DELETE FROM items WHERE ItemID = ?", ItemID)
	if err != nil {
		http.Error(w, "Failed to delete item", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// Get a single item by ID
func GetItem(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    ItemID, _ := strconv.Atoi(params["ItemID"])

    var item models.Item
    err := db.DB.QueryRow("SELECT ItemID, Name, Price, Category  FROM items WHERE ItemID = ?", ItemID).Scan(&item.ItemID, &item.Name, &item.Price, &item.Category)
    if err != nil {
        http.Error(w, "Item not found", http.StatusNotFound)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(item)
}


// Update a item
func UpdateItem(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    ItemID, _ := strconv.Atoi(params["ItemID"])

    var item models.Item
    if err := json.NewDecoder(r.Body).Decode(&item); err != nil {
        http.Error(w, "Invalid input", http.StatusBadRequest)
        return
    }

    _, err := db.DB.Exec("UPDATE items SET Name = ?, Price = ?, Category = ? WHERE ItemID = ?", item.Name, item.Price, item.Category, ItemID)
    if err != nil {
        http.Error(w, "Failed to update item", http.StatusInternalServerError)
        return
    }

    item.ItemID = ItemID
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(item)
}
