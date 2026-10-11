# syntax=docker/dockerfile:1

FROM golang:1.27.1-alpine AS build
WORKDIR /src

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -trimpath -ldflags="-s -w" -o /out/ignicula ./cmd

FROM scratch
COPY --from=build /out/ignicula /ignicula

ENV IGNICULA_ADDR=0.0.0.0
ENV IGNICULA_PORT=16000
EXPOSE 16000

ENTRYPOINT ["/ignicula"]
