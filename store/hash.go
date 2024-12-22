package store

import "sync"

// Hash represents a collection of field-value pairs.
type Hash struct {
	fields map[string]string
	mu     sync.Mutex
}

// NewHash initializes and returns a new Hash.
func NewHash() *Hash {
	return &Hash{fields: make(map[string]string)}
}

// SetField sets a field-value pair in the hash.
func (h *Hash) SetField(field, value string) {
	h.mu.Lock()
	defer h.mu.Unlock()
	h.fields[field] = value
}

// GetField retrieves the value of a specific field.
func (h *Hash) GetField(field string) (string, bool) {
	h.mu.Lock()
	defer h.mu.Unlock()
	value, exists := h.fields[field]
	return value, exists
}
