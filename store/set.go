package store

import "sync"

// Set represents an unordered collection of unique strings.
type Set struct {
	elements map[string]struct{}
	mu       sync.Mutex
}

// NewSet initializes and returns a new Set.
func NewSet() *Set {
	return &Set{elements: make(map[string]struct{})}

}

// Add adds values to the set and returns the count of new items added.
func (s *Set) Add(values ...string) int {
	s.mu.Lock()
	defer s.mu.Unlock()
	count := 0
	for _, value := range values {
		if _, exists := s.elements[value]; !exists {
			s.elements[value] = struct{}{}
			count++
		}
	}
	return count
}

// Members returns all members of the set.
func (s *Set) Members() []string {
	s.mu.Lock()
	defer s.mu.Unlock()

	result := []string{}
	for value := range s.elements {
		result = append(result, value)
	}
	return result
}
