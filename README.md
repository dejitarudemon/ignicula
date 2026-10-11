[![CI](https://github.com/dejitarudemon/ignicula/actions/workflows/ci.yml/badge.svg)](https://github.com/dejitarudemon/ignicula/actions/workflows/ci.yml)
![coverage](https://img.shields.io/endpoint?url=https://gist.githubusercontent.com/dejitarudemon/b9b2b4a363877463d9a2cf0e8f56fd7c/raw/coverage.json)
[![Go Reference](https://pkg.go.dev/badge/github.com/dejitarudemon/ignicula.svg)](https://pkg.go.dev/github.com/dejitarudemon/ignicula)
[![Release](https://img.shields.io/github/v/tag/dejitarudemon/ignicula?label=release)](https://github.com/dejitarudemon/ignicula/releases)
[![License: MIT](https://img.shields.io/github/license/dejitarudemon/ignicula)](LICENSE)

# Ignicula

In-memory Ignicula server on [ignicula-framework](https://github.com/dejitarudemon/ignicula-framework) and [ignicula-wire](https://github.com/dejitarudemon/ignicula-wire).

[English](#english) - [Русский](#русский)

---

## English

> **Not an official release.** Behaviour may still change.  
> No TLS. Data lives in memory only (lost on restart). Authentication is not enforced — Handshake accepts any login.

### Run locally

Go 1.27.1 or newer.

```bash
go run ./cmd
```

Default: `127.0.0.1:16000`. Flags override env; both override defaults. Stop with Ctrl+C.

| Flag | Env | Default |
| --- | --- | --- |
| `-addr` | `IGNICULA_ADDR` | `127.0.0.1` |
| `-port` | `IGNICULA_PORT` | `16000` |

```bash
go run ./cmd -addr 0.0.0.0 -port 16000
IGNICULA_PORT=8080 go run ./cmd
```

### Docker

```bash
docker build -t ignicula .
docker run --rm -p 16000:16000 ignicula
```

The image listens on `0.0.0.0:16000` by default (`IGNICULA_ADDR` / `IGNICULA_PORT` still work).

```bash
docker run --rm -p 8080:8080 -e IGNICULA_PORT=8080 ignicula
```

### Protocol

Binary protocol from [ignicula-wire](https://github.com/dejitarudemon/ignicula-wire): Hello (v0) → Handshake (v1) → Read / Write / Delete / Ping.

Details: [wire tutorial](https://github.com/dejitarudemon/ignicula-wire/blob/main/docs/tutorial.en.md).

---

## Русский

> **Это не официальный релиз.** Поведение ещё может меняться.  
> TLS нет. Данные только в памяти (пропадают при перезапуске). Авторизации нет — Handshake принимает любой login.

### Локальный запуск

Нужен Go 1.27.1 или новее.

```bash
go run ./cmd
```

По умолчанию: `127.0.0.1:16000`. Флаги важнее env; и те и другие важнее дефолтов. Остановка — Ctrl+C.

| Флаг | Env | По умолчанию |
| --- | --- | --- |
| `-addr` | `IGNICULA_ADDR` | `127.0.0.1` |
| `-port` | `IGNICULA_PORT` | `16000` |

```bash
go run ./cmd -addr 0.0.0.0 -port 16000
IGNICULA_PORT=8080 go run ./cmd
```

### Docker

```bash
docker build -t ignicula .
docker run --rm -p 16000:16000 ignicula
```

В образе по умолчанию `0.0.0.0:16000` (`IGNICULA_ADDR` / `IGNICULA_PORT` по-прежнему работают).

```bash
docker run --rm -p 8080:8080 -e IGNICULA_PORT=8080 ignicula
```

### Протокол

Бинарный протокол из [ignicula-wire](https://github.com/dejitarudemon/ignicula-wire): Hello (v0) → Handshake (v1) → Read / Write / Delete / Ping.

Подробности: [туториал wire](https://github.com/dejitarudemon/ignicula-wire/blob/main/docs/tutorial.md).
