package main

import (
    "backend/internal/db"
    "backend/internal/handlers"
    "log"
    "net/http"

    "github.com/gorilla/mux"
)

func main() {
    db.Connect()

    // Set up the router
    r := mux.NewRouter()
    r.HandleFunc("/users", handlers.GetUsers).Methods("GET")
    r.HandleFunc("/users", handlers.CreateUser).Methods("POST")
    r.HandleFunc("/users/{id:[0-9]+}", handlers.GetUser).Methods("GET")
    r.HandleFunc("/users/{id:[0-9]+}", handlers.UpdateUser).Methods("PUT")
    r.HandleFunc("/users/{id:[0-9]+}", handlers.DeleteUser).Methods("DELETE")

    // Start the server
    log.Println("Server is running on port 8080")
    log.Fatal(http.ListenAndServe(":8080", r))
}
