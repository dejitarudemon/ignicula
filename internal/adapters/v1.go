package adapters

import (
	protofields "github.com/dejitarudemon/ignicula-wire/v1/fields"
	protovalue "github.com/dejitarudemon/ignicula-wire/v1/value"
	protovalues "github.com/dejitarudemon/ignicula-wire/v1/value/values"
	"github.com/dejitarudemon/ignicula/internal/value"
)

const (
	opFromV1 = "from v1"
	opToV1   = "to v1"
)

var (
	fromV1 = map[protofields.Type]value.Type{
		protofields.Bytes:        value.Bytes,
		protofields.Float:        value.Float64,
		protofields.Int:          value.Int64,
		protofields.Uint:         value.Uint64,
		protofields.String:       value.String,
		protofields.JSON:         value.JSON,
		protofields.TypedArray:   value.TypedArray,
		protofields.UntypedArray: value.UntypedArray,
	}

	toV1 = map[value.Type]protofields.Type{
		value.Bytes:        protofields.Bytes,
		value.Float64:      protofields.Float,
		value.Int64:        protofields.Int,
		value.Uint64:       protofields.Uint,
		value.String:       protofields.String,
		value.JSON:         protofields.JSON,
		value.TypedArray:   protofields.TypedArray,
		value.UntypedArray: protofields.UntypedArray,
	}
)

// V1Adapter converts between protocol v1 values and storage [value.Value]s.
type V1Adapter struct{}

// FromV1 converts a protocol v1 value into a storage [value.Value].
func (a V1Adapter) FromV1(v protovalue.V) (value.Value, error) {
	if v == nil {
		return value.Value{}, newConversionError(opFromV1, "nil value")
	}

	switch data := v.(type) {
	case protovalues.Int:
		return value.Value{T: value.Int64, V: int64(data)}, nil
	case protovalues.Uint:
		return value.Value{T: value.Uint64, V: uint64(data)}, nil
	case protovalues.Float:
		return value.Value{T: value.Float64, V: float64(data)}, nil
	case protovalues.String:
		return value.Value{T: value.String, V: string(data)}, nil
	case protovalues.JSON:
		return value.Value{T: value.JSON, V: []byte(data)}, nil
	case protovalues.Bytes:
		return value.Value{T: value.Bytes, V: []byte(data)}, nil
	case protovalues.TypedArray:
		return a.fromTypedArray(data)
	case protovalues.UntypedArray:
		return a.fromUntypedArray(data)
	default:
		return value.Value{}, newConversionErrorf(opFromV1, "unsupported wire type %T", v)
	}
}

func (a V1Adapter) fromTypedArray(data protovalues.TypedArray) (value.Value, error) {
	elemType, ok := fromV1[data.ElemType]
	if !ok {
		return value.Value{}, newConversionErrorf(opFromV1, "unsupported typed array element type %s", data.ElemType)
	}

	arr := make([]value.Value, 0, len(data.Elems))
	for i, elem := range data.Elems {
		e, err := a.FromV1(elem)
		if err != nil {
			return value.Value{}, wrapConversionError(opFromV1, "typed array element", err)
		}
		if e.T != elemType {
			return value.Value{}, newConversionErrorf(
				opFromV1,
				"typed array element %d: want %s, got %s",
				i, elemType, e.T,
			)
		}
		arr = append(arr, e)
	}

	return value.Value{
		T: value.TypedArray,
		V: arr,
		M: &value.Metadata{ElemType: elemType},
	}, nil
}

func (a V1Adapter) fromUntypedArray(data protovalues.UntypedArray) (value.Value, error) {
	arr := make([]value.Value, 0, len(data))
	for _, elem := range data {
		e, err := a.FromV1(elem)
		if err != nil {
			return value.Value{}, wrapConversionError(opFromV1, "untyped array element", err)
		}
		arr = append(arr, e)
	}

	return value.Value{T: value.UntypedArray, V: arr}, nil
}

// ToV1 converts a storage [value.Value] into a protocol v1 value.
func (a V1Adapter) ToV1(v value.Value) (protovalue.V, error) {
	switch v.T {
	case value.Int64:
		data, err := payload[int64](v)
		if err != nil {
			return nil, err
		}
		return protovalues.Int(data), nil
	case value.Uint64:
		data, err := payload[uint64](v)
		if err != nil {
			return nil, err
		}
		return protovalues.Uint(data), nil
	case value.Float64:
		data, err := payload[float64](v)
		if err != nil {
			return nil, err
		}
		return protovalues.Float(data), nil
	case value.String:
		data, err := payload[string](v)
		if err != nil {
			return nil, err
		}
		return protovalues.String(data), nil
	case value.JSON:
		data, err := payload[[]byte](v)
		if err != nil {
			return nil, err
		}
		return protovalues.JSON(data), nil
	case value.Bytes:
		data, err := payload[[]byte](v)
		if err != nil {
			return nil, err
		}
		return protovalues.Bytes(data), nil
	case value.TypedArray:
		return a.toTypedArray(v)
	case value.UntypedArray:
		return a.toUntypedArray(v)
	default:
		return nil, newConversionErrorf(opToV1, "unsupported storage type %s", v.T)
	}
}

func (a V1Adapter) toTypedArray(v value.Value) (protovalue.V, error) {
	data, err := payload[[]value.Value](v)
	if err != nil {
		return nil, err
	}
	if v.M == nil {
		return nil, newConversionError(opToV1, "typed array missing element type metadata")
	}

	elemType, ok := toV1[v.M.ElemType]
	if !ok {
		return nil, newConversionErrorf(opToV1, "unsupported typed array element type %s", v.M.ElemType)
	}

	result := protovalues.TypedArray{
		ElemType: elemType,
		Elems:    make([]protovalue.V, 0, len(data)),
	}
	for _, elem := range data {
		e, err := a.ToV1(elem)
		if err != nil {
			return nil, wrapConversionError(opToV1, "typed array element", err)
		}
		result.Elems = append(result.Elems, e)
	}

	return result, nil
}

func (a V1Adapter) toUntypedArray(v value.Value) (protovalue.V, error) {
	data, err := payload[[]value.Value](v)
	if err != nil {
		return nil, err
	}

	result := make(protovalues.UntypedArray, 0, len(data))
	for _, elem := range data {
		e, err := a.ToV1(elem)
		if err != nil {
			return nil, wrapConversionError(opToV1, "untyped array element", err)
		}
		result = append(result, e)
	}

	return result, nil
}

// ToKey converts a protocol key into a storage [value.Key].
func (a V1Adapter) ToKey(key protofields.Key) value.Key {
	return value.Key(key)
}

func payload[T any](v value.Value) (T, error) {
	data, ok := v.V.(T)
	if !ok {
		var zero T
		return zero, UnexpectedPayloadError{Type: v.T, Got: v.V}
	}
	return data, nil
}
