# Use Go 1.23 bookworm as base image
FROM golang:1.25-bookworm AS base

# Move to working directory /build
WORKDIR /build

# Copy the go.mod and go.sum files to the /build directory
COPY go.mod go.sum ./

# Install dependencies
RUN go mod download

# Copy the entire source code into the container
COPY . .

# Build the application
RUN go build -o go-backend

# Document the port that may need to be pudomished
EXPOSE 3000

# Start the application
CMD ["/build/go-backend"]
