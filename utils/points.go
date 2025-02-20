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

// calculates points for a reciept according to the rules
// Rule 1: 1 point for each alphanumeric character in the retailer name.
// Rule 2: 50 points if total is a whole number.
// Rule 3: 25 points if total is a multiple of 0.25.
// Rule 4: 5 points for every two items.
// Rule 5: For each item with a trimmed description length that is a multiple of 3, multiply the price by 0.2 and round up to the nearest integer. 
// Rule 6: 6 points if the purchase day is odd.
// Rule 7: 10 points if the purchase time is between 2:00 PM and 4:00 PM.

func CalculatePoints(receipt models.Receipt) int {
	points := 0

	
	re := regexp.MustCompile(`[a-zA-Z0-9]`)
	points += len(re.FindAllString(receipt.Retailer, -1))
	total, err := strconv.ParseFloat(receipt.Total, 64)
	if err != nil {
		log.Printf("Error parsing total (%s): %v", receipt.Total, err)
		return points
	}

	if total == float64(int(total)) {
		points += 50
	}
	if math.Abs(math.Mod(total, 0.25)) < 1e-9 {
		points += 25
	}

	points += (len(receipt.Items) / 2) * 5

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

	if purchaseDate, err := time.Parse("2006-01-02", receipt.PurchaseDate); err == nil {
		if purchaseDate.Day()%2 == 1 {
			points += 6
		}
	} else {
		log.Printf("Error parsing purchase date (%s): %v", receipt.PurchaseDate, err)
	}

	if purchaseTime, err := time.Parse("15:04", receipt.PurchaseTime); err == nil {
		if purchaseTime.Hour() >= 14 && purchaseTime.Hour() < 16 {
			points += 10
		}
	} else {
		log.Printf("Error parsing purchase time (%s): %v", receipt.PurchaseTime, err)
	}

	return points
}
