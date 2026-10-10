package value

import "fmt"

// Type classifies the payload carried by a [Value].
type Type uint8

const (
	// Bytes is an opaque byte sequence.
	Bytes Type = iota
	// TypedArray is a homogeneous array of values.
	TypedArray
	// UntypedArray is a heterogeneous array of values.
	UntypedArray
	// Int64 is a signed 64-bit integer.
	Int64
	// Uint64 is an unsigned 64-bit integer.
	Uint64
	// Float64 is an IEEE-754 64-bit floating-point number.
	Float64
	// JSON is a JSON-encoded document.
	JSON
)

// String returns a human-readable name for t.
// Unknown values are reported as "Unknown Type (N)".
func (t Type) String() string {
	switch t {
	case Bytes:
		return "Bytes"
	case TypedArray:
		return "Typed Array"
	case UntypedArray:
		return "Untyped Array"
	case Int64:
		return "Int64"
	case Uint64:
		return "Uint64"
	case Float64:
		return "Float64"
	case JSON:
		return "JSON"
	}

	return fmt.Sprintf("Unknown Type (%d)", t)
}
