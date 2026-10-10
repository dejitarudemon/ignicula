// Package value defines keys and typed values stored by Ignicula.
//
// A [Key] identifies an entry. A [Value] pairs a [Type] tag with an arbitrary
// payload in V. Concrete encodings of V for each Type are defined by the
// packages that read and write storage.
package value
