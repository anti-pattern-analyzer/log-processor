# Use an official Golang image as a builder
FROM golang:1.20 as builder

WORKDIR /app

# Copy only go.mod if go.sum is missing
COPY go.mod ./
RUN go mod tidy  # Regenerate go.sum if it's missing

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
