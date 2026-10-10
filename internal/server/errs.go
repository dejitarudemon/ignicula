package server

import "errors"

var (
	// ErrNilStorage is returned by [NewServer] when storage is nil.
	ErrNilStorage = errors.New("server: storage is nil")
)
