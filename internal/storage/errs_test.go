package storage_test

import (
	"errors"
	"testing"

	"github.com/dejitarudemon/ignicula/internal/storage"
	"github.com/dejitarudemon/ignicula/internal/value"
)

func TestNotFoundError(t *testing.T) {
	t.Parallel()

	key := value.Key("missing")
	err := storage.NewNotFoundError(key)

	if got := err.Key(); got != key {
		t.Fatalf("Key() = %q, want %q", got, key)
	}

	wantMsg := "value not found with missing key"
	if got := err.Error(); got != wantMsg {
		t.Fatalf("Error() = %q, want %q", got, wantMsg)
	}

	var notFound storage.NotFoundError
	if !errors.As(err, &notFound) {
		t.Fatal("errors.As(..., &NotFoundError) = false, want true")
	}
}
