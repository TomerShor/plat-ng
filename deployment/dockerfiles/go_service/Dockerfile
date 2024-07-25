FROM golang:1.21-alpine AS build-binary

# Set the working directory
WORKDIR /src

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies
RUN go mod download

# Copy the source code
COPY . .

# Build the binary
RUN go build -o /go/bin/server ./cmd/server

# Command to run the executable
CMD ["/go/bin/server"]