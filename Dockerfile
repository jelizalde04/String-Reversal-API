# Use the official Go image
FROM golang:1.18

# Set the working directory
WORKDIR /app

# Copy the Go module files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod tidy

# Copy the source code
COPY . .

# Expose the port
EXPOSE 8080

# Build and run the Go application
CMD ["go", "run", "main.go"]
