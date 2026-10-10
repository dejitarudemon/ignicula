package adapters

import (
	"fmt"

	"github.com/dejitarudemon/ignicula/internal/value"
)

// ConversionError is returned when a value cannot be converted between the
// wire (v1) representation and the storage [value.Value].
type ConversionError struct {
	op  string
	msg string
	err error
}

func newConversionError(op, msg string) ConversionError {
	return ConversionError{op: op, msg: msg}
}

func newConversionErrorf(op, format string, args ...any) ConversionError {
	return ConversionError{op: op, msg: fmt.Sprintf(format, args...)}
}

func wrapConversionError(op, msg string, err error) ConversionError {
	return ConversionError{op: op, msg: msg, err: err}
}

// Op returns the conversion direction, such as "from v1" or "to v1".
func (e ConversionError) Op() string { return e.op }

// Error implements the error interface.
func (e ConversionError) Error() string {
	if e.err != nil {
		return fmt.Sprintf("adapters: %s: %s: %v", e.op, e.msg, e.err)
	}
	return fmt.Sprintf("adapters: %s: %s", e.op, e.msg)
}

// Unwrap returns the underlying cause, if any.
func (e ConversionError) Unwrap() error { return e.err }

// UnexpectedPayloadError reports that [value.Value.V] does not match [value.Value.T].
type UnexpectedPayloadError struct {
	Type value.Type
	Got  any
}

// Error implements the error interface.
func (e UnexpectedPayloadError) Error() string {
	return fmt.Sprintf("adapters: unexpected payload for %s: %T", e.Type, e.Got)
}
