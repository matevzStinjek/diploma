# Stage 1: Build the application binary
FROM golang:1.24-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /app/worker ./cmd/worker

# Stage 2: Create the final, minimal image
FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/worker .
CMD ["./worker"]
