package main

import (
	"log"
	"net/http"
	"time"

	"receipt-processor/handlers"
	"receipt-processor/store"

	"github.com/gorilla/mux"
)

func main() {
	// Set up the in-memory storage
	memStore := store.NewMemoryStore()

	// Create a new router
	r := mux.NewRouter()

	// Set up handlers with the store
	pHandler := handlers.NewProcessReceipt(memStore)
	pointsHandler := handlers.NewGetPoints(memStore)

	// Map endpoints to handlers
	r.Handle("/receipts/process", pHandler).Methods("POST")
	r.Handle("/receipts/{id}/points", pointsHandler).Methods("GET")

	// Configure the HTTP server
	srv := &http.Server{
		Handler:      r,
		Addr:         ":8080",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Println("Running Receipt Processor on port 8080")
	if err := srv.ListenAndServe(); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
