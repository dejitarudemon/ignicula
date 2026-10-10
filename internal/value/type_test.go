package value_test

import (
	"testing"

	"github.com/dejitarudemon/ignicula/internal/value"
)

func TestTypeString(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		typ  value.Type
		want string
	}{
		{name: "bytes", typ: value.Bytes, want: "Bytes"},
		{name: "typed_array", typ: value.TypedArray, want: "Typed Array"},
		{name: "untyped_array", typ: value.UntypedArray, want: "Untyped Array"},
		{name: "int64", typ: value.Int64, want: "Int64"},
		{name: "uint64", typ: value.Uint64, want: "Uint64"},
		{name: "float64", typ: value.Float64, want: "Float64"},
		{name: "json", typ: value.JSON, want: "JSON"},
		{name: "unknown", typ: value.Type(255), want: "Unknown Type (255)"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			if got := tt.typ.String(); got != tt.want {
				t.Fatalf("Type.String() = %q, want %q", got, tt.want)
			}
		})
	}
}
