# Use the official Go image
FROM golang:1.18

# Set the working directory
WORKDIR /app

# Copy the source code (sin los archivos go.mod y go.sum)
COPY . .

# Run go mod tidy and build the app inside the container
RUN go mod init string-reversal-api
RUN go mod tidy

# Expose the port
EXPOSE 8080

# Build and run the Go application
CMD ["go", "run", "main.go"]
