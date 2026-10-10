// Package syncmap provides an in-memory [storage.Storage] backed by a map
// guarded by a [sync.RWMutex].
package syncmap

import (
	"context"
	"sync"

	"github.com/dejitarudemon/ignicula/internal/storage"
	"github.com/dejitarudemon/ignicula/internal/value"
)

// SyncMap is a concurrency-safe in-memory store.
type SyncMap struct {
	m map[value.Key]value.Value

	mx sync.RWMutex
}

// NewSyncMap returns an empty [SyncMap].
func NewSyncMap() *SyncMap {
	return &SyncMap{
		m: make(map[value.Key]value.Value),
	}
}

// Get returns the value stored under key.
// It returns [storage.NotFoundError] when the key is absent.
func (s *SyncMap) Get(_ context.Context, key value.Key) (value.Value, error) {
	s.mx.RLock()
	defer s.mx.RUnlock()

	v, ok := s.m[key]
	if !ok {
		return value.Value{}, storage.NewNotFoundError(key)
	}

	return v, nil
}

// Put stores value under key, replacing any previous entry.
func (s *SyncMap) Put(_ context.Context, key value.Key, val value.Value) error {
	s.mx.Lock()
	defer s.mx.Unlock()

	s.m[key] = val
	return nil
}

// Delete removes the entry for key. Missing keys are ignored.
func (s *SyncMap) Delete(_ context.Context, key value.Key) error {
	s.mx.Lock()
	defer s.mx.Unlock()

	delete(s.m, key)
	return nil
}
