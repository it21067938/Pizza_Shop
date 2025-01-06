package controller

import (
	"backend/internal/db"
	"backend/internal/models"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Add item to cart
func AddToCart(w http.ResponseWriter, r *http.Request) {
	var cart models.Cart
	if err := json.NewDecoder(r.Body).Decode(&cart); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	_, err := db.DB.Exec("INSERT INTO carts (ItemID, Quantity) VALUES (?, ?)", cart.ItemID, cart.Quantity)
	if err != nil {
		http.Error(w, "Failed to add cart item", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

//Remove item from cart
func RemoveFromCart(w http.ResponseWriter, r *http.Request){
	params := mux.Vars(r)
	CartID, _ := strconv.Atoi(params["CartID"])

	_, err := db.DB.Exec("DELETE FROM carts WHERE CartID = ?", CartID)
	if err != nil {
		http.Error(w, "Failed to delete item", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}


//Get item from cart
func GetCartItems(w http.ResponseWriter, r *http.Request){
	rows, err := db.DB.Query(`
        SELECT carts.ItemID, carts.Quantity, items.Name, items.Price, items.Category
        FROM carts
        JOIN items ON carts.ItemID = items.ItemID
    `)
    if err != nil {
        http.Error(w, "Failed to fetch cart items", http.StatusInternalServerError)
        return
    }
    defer rows.Close()

    var cartItems []models.Cart
    for rows.Next() {
        var cartItem models.Cart
        var item models.Item
        if err := rows.Scan(&cartItem.ItemID, &cartItem.Quantity, &item.Name, &item.Price, &item.Category); err != nil {
            http.Error(w, "Failed to scan cart items", http.StatusInternalServerError)
            return
        }
        cartItem.Item = item
        cartItems = append(cartItems, cartItem)
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(cartItems)
}