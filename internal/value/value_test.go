package value_test

import (
	"testing"

	"github.com/dejitarudemon/ignicula/internal/value"
)

func TestValueHoldsTypeAndPayload(t *testing.T) {
	t.Parallel()

	v := value.Value{
		T: value.Int64,
		V: int64(42),
	}

	if v.T != value.Int64 {
		t.Fatalf("T = %v, want %v", v.T, value.Int64)
	}
	if got, ok := v.V.(int64); !ok || got != 42 {
		t.Fatalf("V = %#v, want int64(42)", v.V)
	}
}

func TestKey(t *testing.T) {
	t.Parallel()

	key := value.Key("user:1")
	if key != "user:1" {
		t.Fatalf("Key = %q, want %q", key, "user:1")
	}
}
