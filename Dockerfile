# Stage 1: Build
FROM golang:1.24.2 AS builder

WORKDIR /app

# Download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Build the binary statically
RUN CGO_ENABLED=0 GOOS=linux go build -o app main.go

# Stage 2: Run
FROM alpine:latest

# Install necessary certs
RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copy the compiled binary and .env file
COPY --from=builder /app/app .
COPY --from=builder /app/.env .env

# Ensure binary is executable
RUN chmod +x ./app

EXPOSE 8080

CMD ["./app"]
