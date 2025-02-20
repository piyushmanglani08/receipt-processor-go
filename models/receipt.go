package models

import (
	"strconv"
	"strings"
	"time"
)

// Receipt represents a purchase receipt.
type Receipt struct {
	Retailer     string `json:"retailer"`
	PurchaseDate string `json:"purchaseDate"`
	PurchaseTime string `json:"purchaseTime"`
	Items        []Item `json:"items"`
	Total        string `json:"total"`
}

// Item represents an item on a receipt.
type Item struct {
	ShortDescription string `json:"shortDescription"`
	Price            string `json:"price"`
}

// Verify checks that required fields are present and valid.
func (r Receipt) Verify() bool {
	if strings.TrimSpace(r.Retailer) == "" ||
		strings.TrimSpace(r.PurchaseDate) == "" ||
		strings.TrimSpace(r.PurchaseTime) == "" ||
		strings.TrimSpace(r.Total) == "" ||
		len(r.Items) == 0 {
		return false
	}
	// Attempt to parse total to ensure it's a valid number.
	if _, err := strconv.ParseFloat(r.Total, 64); err != nil {
		return false
	}
	// Validate purchase date.
	if _, err := time.Parse("2006-01-02", r.PurchaseDate); err != nil {
		return false
	}
	// Validate purchase time.
	if _, err := time.Parse("15:04", r.PurchaseTime); err != nil {
		return false
	}
	// Optionally, add more detailed validation for each item.
	return true
}
