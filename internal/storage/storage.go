package storage

import (
	"context"

	"github.com/dejitarudemon/ignicula/internal/value"
)

type Storage interface {
	Get(ctx context.Context, key value.Key) (value.Value, error)
	Put(ctx context.Context, key value.Key, value value.Value) error
	Delete(ctx context.Context, key value.Key) error
}
