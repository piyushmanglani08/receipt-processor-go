package utils

import (
	"log"
	"math"
	"regexp"
	"strconv"
	"strings"
	"time"

	"receipt-processor-go/models"
)

// CalculatePoints computes points for a receipt based on several rules.
func CalculatePoints(receipt models.Receipt) int {
	points := 0

	// Rule 1: 1 point for each alphanumeric character in the retailer name.
	re := regexp.MustCompile(`[a-zA-Z0-9]`)
	points += len(re.FindAllString(receipt.Retailer, -1))

	// Parse total.
	total, err := strconv.ParseFloat(receipt.Total, 64)
	if err != nil {
		log.Printf("Error parsing total (%s): %v", receipt.Total, err)
		return points
	}

	// Rule 2: 50 points if total is a whole number.
	if total == float64(int(total)) {
		points += 50
	}

	// Rule 3: 25 points if total is a multiple of 0.25.
	if math.Abs(math.Mod(total, 0.25)) < 1e-9 {
		points += 25
	}

	// Rule 4: 5 points for every two items.
	points += (len(receipt.Items) / 2) * 5

	// Rule 5: For each item with a trimmed description length that is a multiple of 3,
	// add points equal to 20% of the item's price (rounded up).
	for _, item := range receipt.Items {
		desc := strings.TrimSpace(item.ShortDescription)
		if len(desc)%3 == 0 {
			price, err := strconv.ParseFloat(item.Price, 64)
			if err != nil {
				log.Printf("Error parsing item price (%s): %v", item.Price, err)
				continue
			}
			points += int(math.Ceil(price * 0.2))
		}
	}

	// Rule 6: 6 points if the purchase day is odd.
	if purchaseDate, err := time.Parse("2006-01-02", receipt.PurchaseDate); err == nil {
		if purchaseDate.Day()%2 == 1 {
			points += 6
		}
	} else {
		log.Printf("Error parsing purchase date (%s): %v", receipt.PurchaseDate, err)
	}

	// Rule 7: 10 points if the purchase time is between 2:00 PM and 4:00 PM.
	if purchaseTime, err := time.Parse("15:04", receipt.PurchaseTime); err == nil {
		if purchaseTime.Hour() >= 14 && purchaseTime.Hour() < 16 {
			points += 10
		}
	} else {
		log.Printf("Error parsing purchase time (%s): %v", receipt.PurchaseTime, err)
	}

	return points
}
