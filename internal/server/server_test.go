package server_test

import (
	"errors"
	"testing"

	"github.com/dejitarudemon/ignicula/internal/server"
	"github.com/dejitarudemon/ignicula/internal/storage/implemented/syncmap"
)

func TestNewServerRequiresStorage(t *testing.T) {
	t.Parallel()

	_, err := server.NewServer(nil, nil)
	if !errors.Is(err, server.ErrNilStorage) {
		t.Fatalf("NewServer(nil, nil) error = %v, want %v", err, server.ErrNilStorage)
	}
}

func TestNewServerOK(t *testing.T) {
	t.Parallel()

	if _, err := server.NewServer(syncmap.NewSyncMap(), nil); err != nil {
		t.Fatalf("NewServer() error = %v", err)
	}
}
