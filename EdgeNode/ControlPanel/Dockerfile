# Use the official Golang image as the base image
FROM golang:1.24.2-alpine

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the rest of the application source code
COPY . .

# Build the Go application
RUN go build -o main .

# Expose port (change if your app uses a different port)
EXPOSE 8080

# Command to run the executable
CMD ["./main"]