# Use an official Go image
FROM golang:1.24.2-alpine

# Set working directory
WORKDIR /app

# Copy Go module files and download deps
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the app
COPY . .

# Build the app
RUN go build -o server .

# Expose the port
EXPOSE 8080

# Run the app
CMD ["./server"]
