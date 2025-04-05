# Use official Golang base image
FROM golang:1.21

# Set the working directory
WORKDIR /app

# Copy go mod files and download deps
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the source code
COPY . .

# Build the Go app
RUN go build -o main .

# Expose port
EXPOSE 8080

# Run the app
CMD ["./main"]
