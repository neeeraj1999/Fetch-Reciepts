package handlers

import (
	"encoding/json"
	"net/http"

	"receipt-processor/models"
	"receipt-processor/store"
	"receipt-processor/utils"

	"github.com/google/uuid"
)

type ProcessReceipt struct {
	store *store.MemoryStore
}

// Initializes the ProcessReceipt handler
func NewProcessReceipt(s *store.MemoryStore) *ProcessReceipt {
	return &ProcessReceipt{store: s}
}

// ServeHTTP handles the POST request to process a receipt
func (pr *ProcessReceipt) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var rcpt models.Receipt
	err := json.NewDecoder(r.Body).Decode(&rcpt)
	if err != nil {
		http.Error(w, "Bad input data. Check your JSON.", http.StatusBadRequest)
		return
	}

	if err := utils.Validate(rcpt); err != nil {
		http.Error(w, "Receipt validation failed. Please check the data.", http.StatusBadRequest)
		return
	}

	newID := uuid.New().String()

	pr.store.InsertReceipt(newID, rcpt)

	resp := models.ReceiptResp{ID: newID}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
