# Build stage
FROM golang:1.21 as builder

WORKDIR /app

# Copy go.mod and go.sum and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code
COPY . .
RUN ls -l /app  # List the files to check if .env is present

# Build the Go application
RUN CGO_ENABLED=0 GOOS=linux go build -o main ./cmd/main.go

# Deployment stage
FROM alpine:latest

WORKDIR /root/
COPY --from=builder /app/main .

# Command to run the executable
CMD ["./main"]
