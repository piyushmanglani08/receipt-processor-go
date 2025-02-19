package storage

import "sync"

var receipts sync.Map

// SaveReceipt stores receipt points
func SaveReceipt(id string, points int) {
    receipts.Store(id, points)
}

// GetPoints retrieves receipt points by ID.
func GetPoints(id string) (int, bool) {
    if val, ok := receipts.Load(id); ok {
        return val.(int), true
    }
    return 0, false
}