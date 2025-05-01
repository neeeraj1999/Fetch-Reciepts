package store

import (
	"sync"

	"receipt-processor/models"
)

// MemoryStore holds receipts in memory
type MemoryStore struct {
	sync.RWMutex
	data map[string]models.Receipt
}

// Creates a new MemoryStore
func NewMemoryStore() *MemoryStore {
	return &MemoryStore{
		data: make(map[string]models.Receipt),
	}
}

// Saves a receipt with its ID
func (ms *MemoryStore) InsertReceipt(id string, rcpt models.Receipt) {
	ms.Lock()
	defer ms.Unlock()
	ms.data[id] = rcpt
}

// Gets a receipt by its ID
func (ms *MemoryStore) FetchReceipt(id string) (models.Receipt, bool) {
	ms.RLock()
	defer ms.RUnlock()
	rcpt, ok := ms.data[id]
	return rcpt, ok
}
