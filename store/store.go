package store

import "sync"

// Store represents the in-memory key-value store with support for multiple data types.
type Store struct {
	data   map[string]interface{}
	mu     sync.RWMutex
	PubSub *PubSub
}

// NewStore initializes and returns a new Store instance.
func NewStore() *Store {
	return &Store{
		data:   make(map[string]interface{}),
		PubSub: NewPubSub(),
	}
}

// Get retrieves a value from the store.
func (s *Store) Get(key string) (string, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	if value, ok := s.data[key].(string); ok {
		return value, true
	}
	return "", false
}

// Set sets a string value in the store.
func (s *Store) Set(key, value string) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.data[key] = value
}

func (s *Store) GetOrCreateList(key string) *List {
	s.mu.Lock()
	defer s.mu.Unlock()
	if l, ok := s.data[key].(*List); ok {
		return l
	}
	l := NewList()
	s.data[key] = l
	return l
}

func (s *Store) GetList(key string) *List {
	s.mu.RLock()
	defer s.mu.RUnlock()
	if l, ok := s.data[key].(*List); ok {
		return l
	}
	return nil
}

func (s *Store) GetOrCreateSet(key string) *Set {
	s.mu.Lock()
	defer s.mu.Unlock()
	if set, ok := s.data[key].(*Set); ok {
		return set
	}
	set := NewSet()
	s.data[key] = set
	return set
}

func (s *Store) GetSet(key string) *Set {
	s.mu.RLock()
	defer s.mu.RUnlock()
	if set, ok := s.data[key].(*Set); ok {
		return set
	}
	return nil
}

func (s *Store) GetOrCreateHash(key string) *Hash {
	s.mu.Lock()
	defer s.mu.Unlock()
	if h, ok := s.data[key].(*Hash); ok {
		return h
	}
	h := NewHash()
	s.data[key] = h
	return h
}

func (s *Store) GetHash(key string) *Hash {
	s.mu.RLock()
	defer s.mu.RUnlock()
	if h, ok := s.data[key].(*Hash); ok {
		return h
	}
	return nil
}
