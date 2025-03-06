# Use an official Golang image as a builder
FROM golang:1.23 as builder

WORKDIR /app

# Copy go.mod and go.sum first to leverage caching
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the source code
COPY . .

# Build the Go application
RUN go build -o go-app

# Use a lightweight image for the final container
FROM alpine:latest
WORKDIR /root/

# Copy the compiled binary from the builder stage
COPY --from=builder /app/go-app .

# Set executable permission and define entrypoint
RUN chmod +x ./go-app
CMD ["./go-app"]
