package main

import (
	"backend/internal/db"
	"backend/internal/routes"
	"github.com/gorilla/handlers"
	"log"
	"net/http"
)

func main() {
	db.Connect()

	r := routes.APIs()

	corsHandler := handlers.CORS(
		handlers.AllowedOrigins([]string{"http://localhost:5173"}),                   // Allow frontend origin
		handlers.AllowedMethods([]string{"GET", "PUT", "POST", "DELETE", "OPTIONS"}), // Allowed HTTP methods
		handlers.AllowedHeaders([]string{"Content-Type", "Authorization"}),           // Allowed headers
	)

	log.Println("Server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", corsHandler(r)))
}
