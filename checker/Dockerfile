# Start from a Debian-based image with the Go 1.16 installed
# and a workspace (GOPATH) configured at /go.
FROM golang:1.19-alpine as builder

# Copy local code to the container image.
WORKDIR /app
COPY . .

# Download dependencies.
RUN go mod download

# Build the command inside the container.
RUN CGO_ENABLED=0 GOOS=linux go build -v -o checker

# Use the official tiny Go image.
# https://hub.docker.com/_/scratch
FROM scratch

# Copy the binary to the production image from the builder stage.
COPY --from=builder /app/checker /checker

# Run the hello binary.
CMD ["/checker"]
