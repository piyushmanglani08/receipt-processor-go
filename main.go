package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"receipt-processor/handlers"
)

func main() {
	router := mux.NewRouter()

	// Define your routes inline
	router.HandleFunc("/receipts/process", handlers.ProcessReceipt).Methods("POST")
	router.HandleFunc("/receipts/{id}/points", handlers.GetPoints).Methods("GET")

	log.Println("Server running on port 8080...")
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
