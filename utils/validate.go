package utils

import (
	"errors"
	"regexp"
	"time"

	"receipt-processor/models"
)

// Checks the receipt for correct structure and data
func Validate(rcpt models.Receipt) error {
	// Check retailer name pattern
	rx := regexp.MustCompile(`^[\w\s\-&]+$`)
	if !rx.MatchString(rcpt.Retailer) {
		return errors.New("retailer name is not valid")
	}

	// Confirm date format
	_, err := time.Parse("2006-01-02", rcpt.Date)
	if err != nil {
		return errors.New("purchase date format is incorrect")
	}

	// Confirm time format
	_, err = time.Parse("15:04", rcpt.Time)
	if err != nil {
		return errors.New("purchase time format is incorrect")
	}

	// Validate total amount format
	totalRx := regexp.MustCompile(`^\d+\.\d{2}$`)
	if !totalRx.MatchString(rcpt.TotalAmount) {
		return errors.New("total amount format is wrong")
	}

	// Ensure there's at least one item
	if len(rcpt.Items) == 0 {
		return errors.New("receipt must contain at least one item")
	}

	// Validate each item
	descRx := regexp.MustCompile(`^[\w\s\-]+$`)
	priceRx := regexp.MustCompile(`^\d+\.\d{2}$`)
	for _, itm := range rcpt.Items {
		if !descRx.MatchString(itm.Description) {
			return errors.New("item description is invalid")
		}
		if !priceRx.MatchString(itm.Price) {
			return errors.New("item price is invalid")
		}
	}

	return nil
}
