# Use official Golang image as base
FROM golang:alpine AS build

# Set working directory
WORKDIR /app

# Copy and download dependencies using go modules
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code into the container
COPY . .

# Build the Go app
RUN go build -o main cmd/main.go

# Production environment
FROM alpine:latest
WORKDIR /root/
COPY --from=build /app/main .
EXPOSE 8080
CMD ["./main"]
