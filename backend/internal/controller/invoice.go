package controller

import (
	"backend/internal/db"
	"backend/internal/models"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

//Add item to invoice
func AddToInvoice(w http.ResponseWriter, r *http.Request){
	var invoice models.Invoice
	var CartData = invoice.CartData

	if !CartData[r.Body.ItemID] {
		CartData[r.Body.ItemID] = 1;
	}else{
		CartData[r.Body.ItemID] += 1;
	}
}