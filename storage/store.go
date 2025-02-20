package storage

import "sync"

type Store struct {
	mu       sync.RWMutex
	receipts map[string]int
}

var store = &Store{
	receipts: make(map[string]int),
}

func SaveReceipt(id string, points int) {
	store.mu.Lock()
	defer store.mu.Unlock()
	store.receipts[id] = points
}

func GetPoints(id string) (int, bool) {
	store.mu.RLock()
	defer store.mu.RUnlock()
	points, exists := store.receipts[id]
	return points, exists
}
