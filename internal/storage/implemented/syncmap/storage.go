package syncmap

import (
	"context"
	"sync"

	"github.com/dejitarudemon/ignicula/internal/storage"
	"github.com/dejitarudemon/ignicula/internal/value"
)

type SyncMap struct {
	m map[*value.Key]value.Value

	mx sync.RWMutex
}

func NewSyncMap() *SyncMap {
	return &SyncMap{
		m: make(map[*value.Key]value.Value),
	}
}

func (s *SyncMap) Get(ctx context.Context, key value.Key) (value.Value, error) {
	s.mx.RLock()
	defer s.mx.RUnlock()

	v, ok := s.m[&key]
	if !ok {
		return value.Value{}, storage.NewNotFoundError(key)
	}

	return v, nil
}

func (s *SyncMap) Put(ctx context.Context, key value.Key, value value.Value) error {
	s.mx.Lock()
	defer s.mx.Unlock()

	s.m[&key] = value
	return nil
}

func (s *SyncMap) Delete(ctx context.Context, key value.Key) error {
	s.mx.Lock()
	defer s.mx.Unlock()

	delete(s.m, &key)
	return nil
}
