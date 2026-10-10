package value

// Value is a typed payload stored under a [Key].
//
// T describes how V should be interpreted. Callers must keep T and V
// consistent; the storage layer does not validate that pairing.
type Value struct {
	T Type
	V any
}
