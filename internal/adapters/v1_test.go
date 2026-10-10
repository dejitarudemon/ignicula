package adapters_test

import (
	"errors"
	"testing"

	protofields "github.com/dejitarudemon/ignicula-wire/v1/fields"
	protovalues "github.com/dejitarudemon/ignicula-wire/v1/value/values"
	"github.com/dejitarudemon/ignicula/internal/adapters"
	"github.com/dejitarudemon/ignicula/internal/value"
)

func TestV1AdapterRoundTripScalars(t *testing.T) {
	t.Parallel()

	a := adapters.V1Adapter{}
	tests := []struct {
		name string
		in   value.Value
	}{
		{name: "int64", in: value.Value{T: value.Int64, V: int64(-7)}},
		{name: "uint64", in: value.Value{T: value.Uint64, V: uint64(7)}},
		{name: "float64", in: value.Value{T: value.Float64, V: 1.5}},
		{name: "string", in: value.Value{T: value.String, V: "hi"}},
		{name: "json", in: value.Value{T: value.JSON, V: []byte(`{"a":1}`)}},
		{name: "bytes", in: value.Value{T: value.Bytes, V: []byte{1, 2, 3}}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			wire, err := a.ToV1(tt.in)
			if err != nil {
				t.Fatalf("ToV1() error = %v", err)
			}
			got, err := a.FromV1(wire)
			if err != nil {
				t.Fatalf("FromV1() error = %v", err)
			}
			if got.T != tt.in.T {
				t.Fatalf("T = %v, want %v", got.T, tt.in.T)
			}
			switch want := tt.in.V.(type) {
			case []byte:
				gotBytes, ok := got.V.([]byte)
				if !ok || string(gotBytes) != string(want) {
					t.Fatalf("V = %#v, want %#v", got.V, tt.in.V)
				}
			default:
				if got.V != tt.in.V {
					t.Fatalf("V = %#v, want %#v", got.V, tt.in.V)
				}
			}
		})
	}
}

func TestV1AdapterTypedArrayRoundTrip(t *testing.T) {
	t.Parallel()

	a := adapters.V1Adapter{}
	in := value.Value{
		T: value.TypedArray,
		V: []value.Value{
			{T: value.Int64, V: int64(1)},
			{T: value.Int64, V: int64(2)},
		},
		M: &value.Metadata{ElemType: value.Int64},
	}

	wire, err := a.ToV1(in)
	if err != nil {
		t.Fatalf("ToV1() error = %v", err)
	}
	ta, ok := wire.(protovalues.TypedArray)
	if !ok {
		t.Fatalf("ToV1() type = %T, want TypedArray", wire)
	}
	if ta.ElemType != protofields.Int {
		t.Fatalf("ElemType = %v, want %v", ta.ElemType, protofields.Int)
	}

	got, err := a.FromV1(wire)
	if err != nil {
		t.Fatalf("FromV1() error = %v", err)
	}
	if got.M == nil || got.M.ElemType != value.Int64 {
		t.Fatalf("metadata = %+v, want ElemType Int64", got.M)
	}
}

func TestV1AdapterToV1UnexpectedPayload(t *testing.T) {
	t.Parallel()

	_, err := adapters.V1Adapter{}.ToV1(value.Value{T: value.Int64, V: "nope"})
	var payloadErr adapters.UnexpectedPayloadError
	if !errors.As(err, &payloadErr) {
		t.Fatalf("error = %v, want UnexpectedPayloadError", err)
	}
}

func TestV1AdapterFromV1Nil(t *testing.T) {
	t.Parallel()

	_, err := adapters.V1Adapter{}.FromV1(nil)
	var conv adapters.ConversionError
	if !errors.As(err, &conv) {
		t.Fatalf("error = %v, want ConversionError", err)
	}
}

func TestV1AdapterToKey(t *testing.T) {
	t.Parallel()

	got := adapters.V1Adapter{}.ToKey(protofields.Key("abc"))
	if got != "abc" {
		t.Fatalf("ToKey() = %q, want %q", got, "abc")
	}
}
