package syncmap_test

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"testing"

	"github.com/dejitarudemon/ignicula/internal/storage"
	"github.com/dejitarudemon/ignicula/internal/storage/implemented/syncmap"
	"github.com/dejitarudemon/ignicula/internal/value"
)

func TestPutGetRoundTrip(t *testing.T) {
	t.Parallel()

	s := syncmap.NewSyncMap()
	ctx := context.Background()
	key := value.Key("answer")
	want := value.Value{T: value.Int64, V: int64(42)}

	if err := s.Put(ctx, key, want); err != nil {
		t.Fatalf("Put() error = %v", err)
	}

	got, err := s.Get(ctx, value.Key("answer"))
	if err != nil {
		t.Fatalf("Get() error = %v", err)
	}
	if got.T != want.T || got.V != want.V {
		t.Fatalf("Get() = %+v, want %+v", got, want)
	}
}

func TestGetMissingReturnsNotFound(t *testing.T) {
	t.Parallel()

	s := syncmap.NewSyncMap()
	_, err := s.Get(context.Background(), value.Key("nope"))
	if err == nil {
		t.Fatal("Get() error = nil, want NotFoundError")
	}

	var notFound storage.NotFoundError
	if !errors.As(err, &notFound) {
		t.Fatalf("Get() error = %v, want NotFoundError", err)
	}
	if notFound.Key() != "nope" {
		t.Fatalf("NotFoundError.Key() = %q, want %q", notFound.Key(), "nope")
	}
}

func TestPutOverwrites(t *testing.T) {
	t.Parallel()

	s := syncmap.NewSyncMap()
	ctx := context.Background()
	key := value.Key("k")

	if err := s.Put(ctx, key, value.Value{T: value.Int64, V: int64(1)}); err != nil {
		t.Fatalf("Put() error = %v", err)
	}
	if err := s.Put(ctx, key, value.Value{T: value.Int64, V: int64(2)}); err != nil {
		t.Fatalf("Put() error = %v", err)
	}

	got, err := s.Get(ctx, key)
	if err != nil {
		t.Fatalf("Get() error = %v", err)
	}
	if got.V != int64(2) {
		t.Fatalf("Get().V = %#v, want int64(2)", got.V)
	}
}

func TestDelete(t *testing.T) {
	t.Parallel()

	s := syncmap.NewSyncMap()
	ctx := context.Background()
	key := value.Key("k")

	if err := s.Put(ctx, key, value.Value{T: value.Bytes, V: []byte("x")}); err != nil {
		t.Fatalf("Put() error = %v", err)
	}
	if err := s.Delete(ctx, value.Key("k")); err != nil {
		t.Fatalf("Delete() error = %v", err)
	}

	_, err := s.Get(ctx, key)
	var notFound storage.NotFoundError
	if !errors.As(err, &notFound) {
		t.Fatalf("Get() after Delete error = %v, want NotFoundError", err)
	}
}

func TestDeleteMissingIsNoOp(t *testing.T) {
	t.Parallel()

	s := syncmap.NewSyncMap()
	if err := s.Delete(context.Background(), value.Key("absent")); err != nil {
		t.Fatalf("Delete() error = %v, want nil", err)
	}
}

func TestConcurrentAccess(t *testing.T) {
	t.Parallel()

	s := syncmap.NewSyncMap()
	ctx := context.Background()

	const workers = 32
	const perWorker = 100

	var wg sync.WaitGroup
	wg.Add(workers)
	for i := range workers {
		go func(id int) {
			defer wg.Done()
			key := value.Key(fmt.Sprintf("%d", id))
			for n := range perWorker {
				_ = s.Put(ctx, key, value.Value{T: value.Int64, V: int64(n)})
				_, _ = s.Get(ctx, key)
				if n%2 == 0 {
					_ = s.Delete(ctx, key)
				}
			}
		}(i)
	}
	wg.Wait()
}
