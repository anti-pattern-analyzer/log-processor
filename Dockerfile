# Use an official Golang image as a builder
FROM golang:1.23 AS builder

WORKDIR /app

# Copy go.mod and go.sum first to leverage caching
COPY go.mod go.sum ./

RUN go mod download

# Copy the rest of the source code
COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o ./go-app ./src

EXPOSE 8085

CMD ["./go-app"]
