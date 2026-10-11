# Changelog

## 0.1.1 - 2026-10-11

Changed:

- Bump ignicula-framework to v0.8.2
- Bump ignicula-wire to v1.2.0

## 0.1.0 - 2026-10-11

First tagged release of the in-memory Ignicula server.

Added:

- TCP server with Hello (v0), Handshake, and v1 Read / Write / Delete / Ping (via ignicula-framework)
- In-memory SyncMap storage and typed value model with v1 wire adapters
- CLI entrypoint (`./cmd`) with `-addr` / `-port` and `IGNICULA_ADDR` / `IGNICULA_PORT`
- Dockerfile (static binary on scratch)
- README (EN/RU)

Notes:

- Not an official protocol release; behaviour may still change
- No TLS, no authentication; all data is in memory only
