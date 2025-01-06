package routes

import (
	"backend/internal/controller"
	"github.com/gorilla/mux"
)

func APIs() *mux.Router {
	r := mux.NewRouter()

	// Item Routes
	r.HandleFunc("/api/food/list", controller.GetItems).Methods("GET")
	r.HandleFunc("/api/food/add", controller.AddItem).Methods("POST")
	r.HandleFunc("/api/food/list/{ItemID:[0-9]+}", controller.GetItem).Methods("GET")
	r.HandleFunc("/api/food/update/{ItemID:[0-9]+}", controller.UpdateItem).Methods("PUT")
	r.HandleFunc("/api/food/remove/{ItemID:[0-9]+}", controller.DeleteItem).Methods("DELETE")

	// Bill Routes
	r.HandleFunc("/api/bills", controller.CreateBill).Methods("POST")
	r.HandleFunc("/api/bills/history", controller.GetAllBills).Methods("GET")

	return r
}