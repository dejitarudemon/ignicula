package value

import "fmt"

type Type uint8

const (
	Bytes Type = iota
	TypedArray
	UntypedArray
	Int64
	Uint64
	Float64
	JSON
)

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
