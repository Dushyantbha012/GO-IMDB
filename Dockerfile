FROM golang:1.20-alpine

# Set environment variables
ENV GO111MODULE=on
ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64

# Set working directory
WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod ./

# Download dependencies
RUN go mod download

# Copy the entire project
COPY . .

# Expose the application's port
EXPOSE 6379

# Run the executable
CMD ["go", "run", "main.go"]