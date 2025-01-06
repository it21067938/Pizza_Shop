package main

import (
    "backend/internal/controller"
    "backend/internal/db"
    "log"
    "net/http"
    "github.com/gorilla/mux"
    "github.com/gorilla/handlers"
)
  

func main() {
    db.Connect()

    r := mux.NewRouter()

    //Item Routes
    r.HandleFunc("/api/food/list", controller.GetItems).Methods("GET")
    r.HandleFunc("/api/food/add", controller.AddItem).Methods("POST")
    r.HandleFunc("/api/food/list/{ItemID:[0-9]+}", controller.GetItem).Methods("GET")
    r.HandleFunc("/api/food/update/{ItemID:[0-9]+}", controller.UpdateItem).Methods("PUT")
    r.HandleFunc("/api/food/remove/{ItemID:[0-9]+}", controller.DeleteItem).Methods("DELETE")

    //Cart routes
    r.HandleFunc("/api/cart", controller.GetCartItems).Methods("GET")
    r.HandleFunc("/api/cart", controller.AddToCart).Methods("POST")
    r.HandleFunc("/api/cart/{CartID:[0-9]+}", controller.RemoveFromCart).Methods("DELETE")

    // Bill Routes
    r.HandleFunc("/api/bill/history", controller.GetAllBills).Methods("GET")

    

    corsHandler := handlers.CORS(
        handlers.AllowedOrigins([]string{"http://localhost:5173"}), // Allow frontend origin
        handlers.AllowedMethods([]string{"GET", "PUT", "POST", "DELETE", "OPTIONS"}), // Allowed HTTP methods
        handlers.AllowedHeaders([]string{"Content-Type", "Authorization"}), // Allowed headers
    )

    log.Println("Server is running on port 8080")
    log.Fatal(http.ListenAndServe(":8080", corsHandler(r)))
}
