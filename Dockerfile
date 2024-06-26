# Builder stage
FROM golang:1.22-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o myserver

# Runner stage
FROM alpine:latest
WORKDIR /app
# Install CA certificates for HTTPS connections
RUN apk --no-cache add ca-certificates
# Copy the built binary from the builder stage
COPY --from=builder /app/myserver /app/myserver
# Expose the port that the app will run on
EXPOSE 8000
# Ensure the binary has execute permissions
RUN chmod +x /app/myserver
# Command to run the binary
CMD ["/app/myserver"]




