package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"receipt-processor/storage"
)

func GetPoints(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	points, exists := storage.GetPoints(id)
	if !exists {
		http.Error(w, "No receipt found for that ID", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(map[string]int{"points": points}); err != nil {
		log.Printf("Failed to encode response for receipt ID %s: %v", id, err)
	}
}
