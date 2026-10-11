package server

import (
	"errors"
	"fmt"

	runtime_v1 "github.com/dejitarudemon/ignicula-framework/runtime/v1"
	"github.com/dejitarudemon/ignicula-wire/v1/fields"
	"github.com/dejitarudemon/ignicula-wire/v1/value"
	"github.com/dejitarudemon/ignicula/internal/storage"
)

func (s *Server) read(ctx runtime_v1.Context, key fields.Key) (value.V, error) {
	s.debug("handling read request", "login", ctx.Login(), "key", key)

	result, err := s.storage.Get(ctx, s.adapterV1.ToKey(key))
	if err != nil {
		if _, ok := errors.AsType[storage.NotFoundError](err); ok {
			s.debug("read request not found", "login", ctx.Login(), "key", key)
			return nil, nil
		}
		s.error("read request failed", "login", ctx.Login(), "key", key, "error", err)
		return nil, fmt.Errorf("server: read %q: %w", key, err)
	}

	out, err := s.adapterV1.ToV1(result)
	if err != nil {
		s.error("read request failed", "login", ctx.Login(), "key", key, "error", err)
		return nil, fmt.Errorf("server: read %q: encode: %w", key, err)
	}

	s.info("read request ok", "login", ctx.Login(), "key", key)
	return out, nil
}

func (s *Server) delete(ctx runtime_v1.Context, key fields.Key) error {
	s.debug("handling delete request", "login", ctx.Login(), "key", key)

	if err := s.storage.Delete(ctx, s.adapterV1.ToKey(key)); err != nil {
		s.error("delete request failed", "login", ctx.Login(), "key", key, "error", err)
		return fmt.Errorf("server: delete %q: %w", key, err)
	}

	s.info("delete request ok", "login", ctx.Login(), "key", key)
	return nil
}

func (s *Server) write(ctx runtime_v1.Context, key fields.Key, v value.V) error {
	s.debug("handling write request", "login", ctx.Login(), "key", key)

	data, err := s.adapterV1.FromV1(v)
	if err != nil {
		s.error("write request failed", "login", ctx.Login(), "key", key, "error", err)
		return fmt.Errorf("server: write %q: decode: %w", key, err)
	}

	if err := s.storage.Put(ctx, s.adapterV1.ToKey(key), data); err != nil {
		s.error("write request failed", "login", ctx.Login(), "key", key, "error", err)
		return fmt.Errorf("server: write %q: %w", key, err)
	}

	s.info("write request ok", "login", ctx.Login(), "key", key)
	return nil
}

func (s *Server) auth(ctx runtime_v1.Context, _ string, _ [32]byte) (bool, error) {
	s.info("auth request ok", "login", ctx.Login())
	return true, nil
}
