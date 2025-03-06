# ---- Build Stage ----
FROM golang:1.23 AS builder
WORKDIR /app

# Copy go.mod and go.sum to leverage caching
COPY go.mod go.sum ./
RUN go mod download

# Copy application source code
COPY . .

# Build the application (static binary, no dependencies)
RUN CGO_ENABLED=0 GOOS=linux go build -o go-app .

# ---- Run Stage (Minimal) ----
FROM alpine:latest
WORKDIR /app

# Copy only the compiled binary from the builder stage
COPY --from=builder /app/go-app .

# Set permissions and expose the port
RUN chmod +x go-app
EXPOSE 8085

# Run the application
CMD ["./go-app"]
