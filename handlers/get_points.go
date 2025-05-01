package handlers

import (
	"encoding/json"
	"net/http"

	"receipt-processor/models"
	"receipt-processor/store"
	"receipt-processor/utils"

	"github.com/gorilla/mux"
)

type GetPoints struct {
	store *store.MemoryStore
}

// Initialize the GetPoints handler
func NewGetPoints(s *store.MemoryStore) *GetPoints {
	return &GetPoints{store: s}
}

// ServeHTTP processes the GET request for points
func (gp *GetPoints) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	receiptID := params["id"]

	rcpt, found := gp.store.FetchReceipt(receiptID)
	if !found {
		http.Error(w, "No such receipt.", http.StatusNotFound)
		return
	}

	totalPoints, err := utils.CalcPoints(rcpt)
	if err != nil {
		http.Error(w, "Failed to compute points.", http.StatusInternalServerError)
		return
	}

	resp := models.PointsResp{Points: totalPoints}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
