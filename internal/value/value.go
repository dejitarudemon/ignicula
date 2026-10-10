package value

// Metadata carries optional type details that do not fit in [Value.T] alone.
type Metadata struct {
	// ElemType is the homogeneous element type of a [TypedArray] value.
	ElemType Type
}

// Value is a typed payload stored under a [Key].
//
// T describes how V should be interpreted. Callers must keep T and V
// consistent; the storage layer does not validate that pairing.
type Value struct {
	T Type
	V any
	M *Metadata
}
