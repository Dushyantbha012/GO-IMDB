package store

import "sync"

// List represents a list data type (FIFO queue).
type List struct {
	elements []string
	mu       sync.Mutex
}

// NewList initializes and returns a new List.
func NewList() *List {
	return &List{elements: []string{}}
}

// LPush prepends values to the list.
func (l *List) LPush(values ...string) int {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.elements = append(values, l.elements...)
	return len(l.elements)
}

// RPush appends values to the list.
func (l *List) RPush(values ...string) int {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.elements = append(l.elements, values...)
	return len(l.elements)
}

// LPop removes and returns the first element.
func (l *List) LPop() (string, bool) {
	l.mu.Lock()
	defer l.mu.Unlock()
	if len(l.elements) == 0 {
		return "", false
	}
	value := l.elements[0]
	l.elements = l.elements[1:]
	return value, true
}

// RPop removes and returns the last element.
func (l *List) RPop() (string, bool) {
	l.mu.Lock()
	defer l.mu.Unlock()
	if len(l.elements) == 0 {
		return "", false
	}
	value := l.elements[len(l.elements)-1]
	l.elements = l.elements[:len(l.elements)-1]
	return value, true
}
