# ──────────────────────────────
# 1. Build stage
# ──────────────────────────────
FROM golang:1.24-alpine AS builder

# Собираем статически, для Linux/amd64
ENV CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /src

# Кэшируем зависимости
COPY go.mod go.sum ./
RUN --mount=type=cache,target=/go/pkg/mod \
    go mod download

# Копируем остальной код и собираем
COPY . .

ARG BUILD_TARGET=server

RUN --mount=type=cache,target=/go/pkg/mod \
    --mount=type=cache,target=/root/.cache/go-build \
    go build -v -o /bin/app ./cmd/$BUILD_TARGET

# ──────────────────────────────
# 2. Runtime stage
# ──────────────────────────────
FROM debian:12-slim AS runtime

RUN apt-get update

COPY --from=builder /bin/app /app

ENTRYPOINT ["/app"]