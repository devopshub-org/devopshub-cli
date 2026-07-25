# syntax=docker/dockerfile:1
# ─────────────────────────────────────────────────────────────────
# Stage 1: Builder
# ─────────────────────────────────────────────────────────────────
FROM golang:1.22-alpine AS builder

ARG VERSION=dev
ARG COMMIT=none
ARG DATE=unknown
ARG BUILT_BY=docker

WORKDIR /src

# Download dependencies first — leverages Docker layer cache.
COPY go.mod go.sum ./
RUN go mod download && go mod verify

# Copy source and build.
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build \
      -trimpath \
      -ldflags="-s -w \
        -X github.com/devopshub-org/devopshub-cli/internal/build.Version=${VERSION} \
        -X github.com/devopshub-org/devopshub-cli/internal/build.Commit=${COMMIT} \
        -X github.com/devopshub-org/devopshub-cli/internal/build.Date=${DATE} \
        -X github.com/devopshub-org/devopshub-cli/internal/build.BuiltBy=${BUILT_BY}" \
      -o /devopshub \
      ./cmd/devopshub

# ─────────────────────────────────────────────────────────────────
# Stage 2: Runtime (distroless — minimal attack surface)
# ─────────────────────────────────────────────────────────────────
FROM gcr.io/distroless/static-debian12:nonroot

LABEL org.opencontainers.image.title="devopshub-cli"
LABEL org.opencontainers.image.description="The open-source DevOps toolkit for engineers"
LABEL org.opencontainers.image.url="https://github.com/devopshub-org/devopshub-cli"
LABEL org.opencontainers.image.source="https://github.com/devopshub-org/devopshub-cli"
LABEL org.opencontainers.image.licenses="Apache-2.0"

COPY --from=builder /devopshub /devopshub

USER nonroot:nonroot
ENTRYPOINT ["/devopshub"]
