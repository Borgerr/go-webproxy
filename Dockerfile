# Use Go 1.24 as base image
FROM golang:1.24 AS base

# move to working directory /build
WORKDIR /build
COPY go.mod go.sum ./

# Install dependencies
RUN go mod download

# Copy entire source code into container
COPY . .

# Specify port- change here for using other ports
ARG PORT
ENV PORT=2100

# Document that port may need to be published by default
EXPOSE $PORT

# Build the application
RUN go build

# Start the application
CMD ["sh", "-c", "/build/go-webproxy --port $PORT"]
