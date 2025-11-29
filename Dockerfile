# syntax=docker/dockerfile:1

# Builder stage
FROM golang:1.25-bookworm AS builder

WORKDIR /app

# Cache dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy source
COPY . .

# Provide a default config file inside the image
RUN cp configs/configs.example.yaml configs/configs.yaml

# Build the server binary
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /app/bin/seat-reservation ./cmd/server

# Final image
FROM gcr.io/distroless/static-debian12:nonroot

WORKDIR /app

# Copy binary and default config
COPY --from=builder /app/bin/seat-reservation /app/seat-reservation
COPY --from=builder /app/configs /app/configs

# Application listens on port 8080 by default (configurable via config file)
EXPOSE 8080

ENTRYPOINT ["/app/seat-reservation"]