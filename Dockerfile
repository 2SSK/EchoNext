# Multi-stage build for Go backend with embedded Next.js static files
FROM node:20-alpine AS frontend-builder

WORKDIR /frontend

# Copy frontend package files
COPY app/package.json app/pnpm-lock.yaml ./

# Install pnpm and dependencies
RUN npm install -g pnpm && pnpm install --frozen-lockfile

# Copy frontend source
COPY app/ .

# Build frontend
RUN pnpm build

# Go build stage
FROM golang:1.25-alpine AS backend-builder

WORKDIR /app

# Copy Go modules
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy source code
COPY . .

# Copy built frontend static files
COPY --from=frontend-builder /frontend/out ./app/out

# Build Go binary
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main cmd/EchoNext/main.go

# Final stage
FROM alpine:latest

# Install ca-certificates for HTTPS
RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copy binary
COPY --from=backend-builder /app/main .

# Copy static files
COPY --from=backend-builder /app/app/out ./app/out

# Expose port
EXPOSE 8080

# Run
CMD ["./main"]