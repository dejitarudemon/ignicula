package server

import (
	"fmt"

	"github.com/dejitarudemon/ignicula-framework/logger"
	runtime_v1 "github.com/dejitarudemon/ignicula-framework/runtime/v1"
	rconfig "github.com/dejitarudemon/ignicula-framework/runtime/v1/config"
	fwserver "github.com/dejitarudemon/ignicula-framework/server"
	sconfig "github.com/dejitarudemon/ignicula-framework/server/config"
	"github.com/dejitarudemon/ignicula/internal/adapters"
	"github.com/dejitarudemon/ignicula/internal/storage"
)

// Server is an Ignicula node: framework transport plus application storage.
type Server struct {
	inner     *fwserver.Server
	storage   storage.Storage
	logger    logger.Logger
	adapterV1 adapters.V1Adapter
}

// NewServer builds a server that serves storage over the v1 protocol runtime.
// logger may be nil; storage must be non-nil.
func NewServer(store storage.Storage, log logger.Logger) (*Server, error) {
	if store == nil {
		return nil, ErrNilStorage
	}

	s := &Server{
		inner:     fwserver.NewServer(sconfig.NewServerConfig(), log),
		storage:   store,
		logger:    log,
		adapterV1: adapters.V1Adapter{},
	}

	runtime := runtime_v1.NewRuntimeBuilder(*rconfig.NewRuntimeBuilderConfig()).
		WithHandlerRead(s.read).
		WithHandlerWrite(s.write).
		WithHandlerDelete(s.delete).
		WithHandlerAuth(s.auth).
		Build()

	if err := s.inner.RegisterV1Runtime(&runtime); err != nil {
		return nil, fmt.Errorf("server: register v1 runtime: %w", err)
	}

	return s, nil
}

// Close stops the server and releases its resources.
func (s *Server) Close() error {
	if err := s.inner.Close(); err != nil {
		return fmt.Errorf("server: close: %w", err)
	}
	return nil
}

// Start begins listening on addr:port.
func (s *Server) Start(addr string, port uint16) error {
	if err := s.inner.Start(addr, port); err != nil {
		return fmt.Errorf("server: start %s:%d: %w", addr, port, err)
	}
	return nil
}
