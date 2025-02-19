# === BUILD STAGE ===
# Use the official Golang image to build the application
FROM golang:1.24 AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files to the container to install dependencies
COPY go.mod go.sum ./

# Download Go dependencies (caching them to speed up future builds)
RUN go mod download

# Copy the source code into the container
COPY . .

# Build the Go application into a binary named 'main'
RUN go build -o main .

# === FINAL STAGE ===
# Start a new stage from a smaller image to keep the final image lightweight
FROM gcr.io/distroless/base

# Set the working directory for the runtime environment
WORKDIR /root/

# Copy the compiled binary from the builder stage
COPY --from=builder /app/main .

# Expose port 8080 to allow the app to be accessed externally
EXPOSE 8080

# Define the command to run the application when the container starts
CMD ["./main"]
