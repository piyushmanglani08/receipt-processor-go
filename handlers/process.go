package handlers

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/google/uuid"
	"receipt-processor-go/models"
	"receipt-processor-go/storage"
	"receipt-processor-go/utils"
)

// ProcessReceipt decodes a receipt, validates it, calculates its points, and stores the result.
func ProcessReceipt(w http.ResponseWriter, r *http.Request) {
	// Read the raw body for deterministic ID generation.
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Printf("Error reading receipt body: %v", err)
		utils.WriteError(w, "The receipt is invalid.", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	var receipt models.Receipt
	if err := json.Unmarshal(body, &receipt); err != nil {
		log.Printf("Error decoding receipt: %v", err)
		utils.WriteError(w, "The receipt is invalid.", http.StatusBadRequest)
		return
	}

	// Validate receipt fields.
	if !receipt.Verify() {
		log.Printf("Receipt verification failed: %+v", receipt)
		utils.WriteError(w, "The receipt is invalid.", http.StatusBadRequest)
		return
	}

	// Generate a deterministic ID using SHA-1.
	id := uuid.NewSHA1(uuid.Max, body).String()

	// Calculate points.
	points := utils.CalculatePoints(receipt)

	// Save receipt data.
	storage.SaveReceipt(id, points)
	utils.WriteJSON(w, map[string]string{"id": id})
}
