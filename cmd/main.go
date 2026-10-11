package main

import (
	"flag"
	"fmt"
	"log/slog"
	"os"
	"os/signal"
	"strconv"

	logimpl "github.com/dejitarudemon/ignicula/internal/logger/implemented"
	"github.com/dejitarudemon/ignicula/internal/server"
	"github.com/dejitarudemon/ignicula/internal/storage/implemented/syncmap"
)

const (
	defaultAddr = "127.0.0.1"
	defaultPort = 16000

	envAddr = "IGNICULA_ADDR"
	envPort = "IGNICULA_PORT"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "ignicula: %v\n", err)
		os.Exit(1)
	}
}

func run() error {
	addrFlag := flag.String("addr", "", "listen IP/host (env "+envAddr+", default "+defaultAddr+")")
	portFlag := flag.Uint("port", 0, "listen port (env "+envPort+", default "+strconv.Itoa(defaultPort)+")")
	flag.Parse()

	addr, err := resolveAddr(*addrFlag, os.Getenv(envAddr))
	if err != nil {
		return err
	}
	port, err := resolvePort(*portFlag, os.Getenv(envPort))
	if err != nil {
		return err
	}

	log := logimpl.New(slog.Default())
	srv, err := server.NewServer(syncmap.NewSyncMap(), log)
	if err != nil {
		return err
	}

	if err := srv.Start(addr, port); err != nil {
		return err
	}
	log.Info("listening", "addr", addr, "port", port)

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)
	<-ch
	log.Info("shutting down")

	return srv.Close()
}

func resolveAddr(flagValue, envValue string) (string, error) {
	if flagValue != "" {
		return flagValue, nil
	}
	if envValue != "" {
		return envValue, nil
	}
	return defaultAddr, nil
}

func resolvePort(flagValue uint, envValue string) (uint16, error) {
	if flagValue != 0 {
		if flagValue > 65535 {
			return 0, fmt.Errorf("port %d out of range", flagValue)
		}
		return uint16(flagValue), nil
	}
	if envValue != "" {
		n, err := strconv.ParseUint(envValue, 10, 16)
		if err != nil {
			return 0, fmt.Errorf("%s: %w", envPort, err)
		}
		if n == 0 {
			return 0, fmt.Errorf("%s: port must be 1-65535", envPort)
		}
		return uint16(n), nil
	}
	return defaultPort, nil
}
