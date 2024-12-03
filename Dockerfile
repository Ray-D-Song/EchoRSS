# Build frontend
FROM node:20-alpine AS frontend-builder
WORKDIR /app
# Install pnpm
RUN npm install -g pnpm

# Copy frontend files
COPY web/package.json web/pnpm-lock.yaml ./
RUN pnpm install

COPY web/ .
RUN pnpm run build

# Build backend
FROM golang:1.23-alpine AS backend-builder
WORKDIR /app
COPY server/ .
COPY --from=frontend-builder /app/dist ./dist
RUN go mod tidy
RUN go build -o echo-rss

# Final image
FROM alpine:latest
WORKDIR /app

# Install necessary runtime dependencies
RUN apk add --no-cache ca-certificates

# Copy build artifacts
COPY --from=backend-builder /app/echo-rss .
COPY --from=backend-builder /app/db/migrations ./db/migrations

# Create necessary directories
RUN mkdir -p resources/logs

# Expose port
EXPOSE 11299

# Start application
CMD ["./echo-rss"]
