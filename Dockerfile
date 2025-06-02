FROM golang:1.24.2-alpine AS builder

WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the source code
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -o psygrow-api ./src/cmd/api

# Use a smaller image for the final image
FROM alpine:latest

WORKDIR /app

# Copy the binary from the builder stage
COPY --from=builder /app/psygrow-api .

# Expose the port the app runs on
EXPOSE 8080

# Command to run the application
CMD ["./psygrow-api"]