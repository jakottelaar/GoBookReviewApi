# Stage 1: Build the Go app
FROM golang:1.22.4 AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files to the working directory
COPY go.mod go.sum ./

# Download the Go module dependencies
RUN go mod download

# Copy the entire project directory to the working directory inside the container
COPY . .

# Build the Go app with CGO disabled for a static binary
# Make sure to point to the directory where main.go is located, e.g., ./cmd/api
RUN CGO_ENABLED=0 GOOS=linux go build -o /docker-book-review-app ./cmd/api

# Stage 2: Create a smaller runtime image
FROM alpine:latest

# Install the CA certificates
RUN apk --no-cache add ca-certificates

# Set the working directory inside the container
WORKDIR /root/

# Copy the binary from the builder stage
COPY --from=builder /docker-book-review-app .


# Command to run the executable
CMD ["./docker-book-review-app"]
