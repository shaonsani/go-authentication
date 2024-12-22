# Use a newer Go version that satisfies the requirement
FROM golang:1.23.4-alpine

# Set environment variables
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 \
    PATH="/go/bin:$PATH"

# Set working directory
WORKDIR /app

# Install dependencies for Alpine Linux
RUN apk add --no-cache git postgresql-client

# Initialize the Go module
COPY go.mod go.sum ./
RUN [ -f go.mod ] || go mod init example.com/app
RUN go mod tidy

# Install swag CLI for Swagger documentation generation
RUN go install github.com/swaggo/swag/cmd/swag@latest

# Verify that swag is installed correctly
RUN swag --version

# Copy the entire application code
COPY . .

# Generate Swagger documentation
RUN swag init

# Copy the .env file into the container
COPY .env .env

# Build the application
RUN go build -o main .

# Expose the application port
EXPOSE 8080

# Set the entry point for the container
CMD ["./main"]
