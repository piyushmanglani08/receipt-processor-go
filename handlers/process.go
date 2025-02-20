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

// checks if recipt is good to be sent for calculation of points 
func ProcessReceipt(w http.ResponseWriter, r *http.Request) {
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


	if !receipt.Verify() {
		log.Printf("Receipt verification failed: %+v", receipt)
		utils.WriteError(w, "The receipt is invalid.", http.StatusBadRequest)
		return
	}

	id := uuid.NewSHA1(uuid.Max, body).String()

	points := utils.CalculatePoints(receipt)

	storage.SaveReceipt(id, points)
	utils.WriteJSON(w, map[string]string{"id": id})
}
