# Build frontend
FROM node:20-alpine AS frontend-builder
RUN npm install -g pnpm

# Copy frontend files
WORKDIR /web
COPY web/package.json web/pnpm-lock.yaml ./
RUN pnpm install

COPY web/ .
RUN pnpm run build

# Build backend
FROM golang:1.23-alpine AS backend-builder
WORKDIR /server
RUN apk add --no-cache gcc musl-dev

COPY server/ .
COPY --from=frontend-builder /server/dist ./dist
RUN go mod tidy
RUN CGO_ENABLED=1 go build -o echo-rss

# Final image
FROM alpine:latest

# Copy build artifacts
COPY --from=backend-builder /server/echo-rss .
COPY --from=backend-builder /server/dist ./dist

# Expose port
EXPOSE 11299

# Start application
CMD ["./echo-rss"]
