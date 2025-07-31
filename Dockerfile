# ─── Stage 1: Build ───────────────────────────────────────
FROM golang:1.24.5-alpine AS builder

# Install git for module fetching
RUN apk add --no-cache git

# Set working directory
WORKDIR /app

# Cache go mod
COPY go.mod ./
COPY go.sum ./
RUN go mod download

# Copy source and build
COPY . ./
RUN go build -o shinyvanity ./cmd/shinyvanity

# ─── Stage 2: Runtime ─────────────────────────────────────
FROM alpine:latest

# Set non-root user for security
RUN adduser -D -g '' appuser
USER appuser

WORKDIR /home/appuser

# Copy binary from builder
COPY --from=builder /app/shinyvanity .

# Expose port (adjust if needed)
EXPOSE 8080

# Run the binary
ENTRYPOINT ["./shinyvanity"]
