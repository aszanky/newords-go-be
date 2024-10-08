FROM golang:1.20-alpine AS builder

WORKDIR /app

ENV GO111MODULE=on

# Copy go mod and sum files
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the application files
COPY . .

# Copy .env file
COPY .env .env

# Build the Go application
RUN go build -o /newords-go-be

EXPOSE 8090

# Use a minimal image to run the Go application
# Multi stage build 
FROM scratch

COPY --from=builder /newords-go-be /newords-go-be

ENTRYPOINT ["/newords-go-be"]
