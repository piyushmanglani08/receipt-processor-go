package handlers

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"receipt-processor-go/storage"
	"receipt-processor-go/utils"
)

// GetPoints retrieves the points for a given receipt ID.
func GetPoints(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	points, exists := storage.GetPoints(id)
	if !exists {
		utils.WriteError(w, "No receipt found for that ID", http.StatusNotFound)
		return
	}

	if err := utils.WriteJSON(w, map[string]int{"points": points}); err != nil {
		log.Printf("Failed to encode response for receipt ID %s: %v", id, err)
	}
}
