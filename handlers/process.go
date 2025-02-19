package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/google/uuid"
	"receipt-processor/models"
	"receipt-processor/storage"
	"receipt-processor/utils"
)


func ProcessReceipt(w http.ResponseWriter, r *http.Request) {
	var receipt models.Receipt

	if err := json.NewDecoder(r.Body).Decode(&receipt); err != nil {
		log.Printf("Error decoding receipt: %v", err)
		http.Error(w, "The receipt is invalid.", http.StatusBadRequest)
		return
	}

	id := uuid.New().String()

	points := utils.CalculatePoints(receipt)

	storage.SaveReceipt(id, points)

	response := map[string]string{"id": id}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		log.Printf("Error encoding response: %v", err)
	}
}
