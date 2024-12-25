# Build stage: Use the official Go image to compile the application
FROM --platform=$BUILDPLATFORM golang:1.23 AS build

# Set the Current Working Directory inside the container
WORKDIR /cryptonian

# Copy go.mod and go.sum files
COPY go.mod .
COPY go.sum .

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

ADD . .

WORKDIR /cryptonian/cmd/server
# Build the Go app. Adjust the path to main.go accordingly
RUN go build -o server

# Final stage: Use a minimal Debian-based image to run the application
FROM debian:bookworm-slim

# Install the ca-certificates package
RUN apt-get update && \
    apt-get install -y ca-certificates curl && \
    rm -rf /var/lib/apt/lists/* && \
    mkdir cryptonian

# Set the Current Working Directory inside the container
WORKDIR /cryptonian

ENV CONFIG_FILE=/cryptonian/prod.yaml

# Copy the pre-built binary file from the previous stage
COPY --from=build /cryptonian/cmd/server/server .
COPY --from=build /cryptonian/cmd/config/prod.yaml ./prod.yaml

# create a directory for the logs
RUN mkdir temp

# Expose port 8080 to the outside world
EXPOSE 8080

HEALTHCHECK --interval=30s --timeout=1s --start-period=60s --retries=3 CMD curl -f http://localhost:8080/health || exit 1

# Command to run the executable
CMD ["./server"]