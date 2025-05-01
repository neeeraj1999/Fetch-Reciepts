package utils

import (
	"math"
	"regexp"
	"strconv"
	"strings"
	"time"

	"receipt-processor/models"
)

// Calculates the points based on receipt details
func CalcPoints(rcpt models.Receipt) (int, error) {
	points := 0

	// Points for alphanumerics in retailer name
	alphaRegex := regexp.MustCompile(`[A-Za-z0-9]`)
	alphaCount := len(alphaRegex.FindAllString(rcpt.Retailer, -1))
	points += alphaCount

	// Bonus for whole dollar total
	totalVal, err := strconv.ParseFloat(rcpt.TotalAmount, 64)
	if err != nil {
		return 0, err
	}
	if totalVal == float64(int(totalVal)) {
		points += 50
	}

	// Bonus for total divisible by 0.25
	if math.Mod(totalVal, 0.25) == 0 {
		points += 25
	}

	// Points for every two items
	itemCount := len(rcpt.Items)
	points += (itemCount / 2) * 5

	// Points for item descriptions
	for _, itm := range rcpt.Items {
		desc := strings.TrimSpace(itm.Description)
		if len(desc)%3 == 0 {
			price, err := strconv.ParseFloat(itm.Price, 64)
			if err != nil {
				return 0, err
			}
			pt := int(math.Ceil(price * 0.2))
			points += pt
		}
	}

	// Bonus for odd purchase day
	pDate, err := time.Parse("2006-01-02", rcpt.Date)
	if err != nil {
		return 0, err
	}
	if pDate.Day()%2 != 0 {
		points += 6
	}

	// Bonus for purchase time between 2pm and 4pm
	pTime, err := time.Parse("15:04", rcpt.Time)
	if err != nil {
		return 0, err
	}
	totalMinutes := pTime.Hour()*60 + pTime.Minute()
	if totalMinutes > 14*60 && totalMinutes < 16*60 {
		points += 10
	}

	return points, nil
}
