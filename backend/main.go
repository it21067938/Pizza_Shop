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
    r.HandleFunc("/api/food/list", controller.GetItems).Methods("GET")
    r.HandleFunc("/api/food/add", controller.AddItem).Methods("POST")
    r.HandleFunc("/api/food/list/{ItemID:[0-9]+}", controller.GetItem).Methods("GET")
    r.HandleFunc("/api/food/update/{ItemID:[0-9]+}", controller.UpdateItem).Methods("PUT")
    r.HandleFunc("/api/food/remove/{ItemID:[0-9]+}", controller.DeleteItem).Methods("DELETE")

    corsHandler := handlers.CORS(
        handlers.AllowedOrigins([]string{"http://localhost:5173"}), // Allow frontend origin
        handlers.AllowedMethods([]string{"GET", "PUT", "POST", "DELETE", "OPTIONS"}), // Allowed HTTP methods
        handlers.AllowedHeaders([]string{"Content-Type", "Authorization"}), // Allowed headers
    )

    log.Println("Server is running on port 8080")
    log.Fatal(http.ListenAndServe(":8080", corsHandler(r)))
}
