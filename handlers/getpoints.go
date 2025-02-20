package handlers

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"receipt-processor-go/storage"
	"receipt-processor-go/utils"
)

// getting points for a given id
func GetPoints(w http.ResponseWriter, r *http.Request) {

    id := mux.Vars(r)["id"]
    points, exists := storage.GetPoints(id)
    if !exists {
        utils.WriteError(w, fmt.Sprintf("No receipt found for ID: %s", id), http.StatusNotFound)
        return
    }
    
    response := map[string]int{"points": points}
    
    if err := utils.WriteJSON(w, response); err != nil {
        log.Printf("Error encoding JSON response for receipt ID %s: %v", id, err)
        utils.WriteError(w, "Internal server error", http.StatusInternalServerError)
    }
}
