package cache

import (
	"log"
	"sync"
	"time"

	"github.com/abvarun226/ristretto-cache/ds"
	"github.com/dgraph-io/ristretto"
)

// New initializes a new cache store.
func New() *Store {
	c, errNew := ristretto.NewCache(&ristretto.Config{
		MaxCost:     1 << 30,
		NumCounters: 1e7,
		BufferItems: 1,
	})
	if errNew != nil {
		log.Fatalf("failed to initialize cache: %v", errNew)
	}
	return &Store{store: c}
}

// Store is our struct for cache store.
type Store struct {
	l     sync.RWMutex
	store *ristretto.Cache
}

// Get method to retrieve the value of a key. If not present, returns false.
func (s *Store) Get(key string) (string, bool) {
	s.l.RLock()
	defer s.l.RUnlock()

	var value string
	val, found := s.store.Get(key)
	if found {
		value = val.(string)
	}
	return value, found
}

// Delete method to delete a key from cahce.
func (s *Store) Delete(key string) {
	s.l.Lock()
	defer s.l.Unlock()

	s.store.Del(key)
}

// SetByTags method to set cache by given tags.
func (s *Store) SetByTags(key, value string, expiry time.Duration, tags []string) {
	s.l.Lock()
	defer s.l.Unlock()

	for _, tag := range tags {
		set := ds.New()
		if v, found := s.store.Get(tag); found {
			set = v.(*ds.StringSet)
		}
		set.Add(key)
		s.store.Set(tag, set, 1)
	}

	s.store.SetWithTTL(key, value, 1, expiry)
}

// Invalidate method to invalidate cache with given tags.
func (s *Store) Invalidate(tags []string) {
	s.l.Lock()
	defer s.l.Unlock()

	keys := make([]string, 0)
	for _, tag := range tags {
		set := ds.New()
		if v, found := s.store.Get(tag); found {
			set = v.(*ds.StringSet)
		}
		keys = append(keys, set.Members()...)
		keys = append(keys, tag)
	}

	for _, k := range keys {
		s.store.Del(k)
	}
}

// Close method clear and then close the cache store.
func (s *Store) Close() {
	s.store.Clear()
	s.store.Close()
}
