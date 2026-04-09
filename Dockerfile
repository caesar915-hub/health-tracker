# STAGE 1: Build the binary
FROM golang:alpine AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy the go.mod and go.sum files first (for caching layers)
COPY backend/go.mod backend/go.sum ./
RUN go mod download

# Copy the actual source code
COPY backend/ .

# Build the app into a static binary called "server"
RUN CGO_ENABLED=0 GOOS=linux go build -o /server .

# STAGE 2: Run the binary
FROM alpine:latest

WORKDIR /

# Copy the binary from the builder stage
COPY --from=builder /server /server

# Copy your frontend folder (The binary needs this to serve index.html!)
COPY frontend /frontend

# Tell Docker to listen on port 8080
EXPOSE 8080

# Run the app
CMD ["/server"]
