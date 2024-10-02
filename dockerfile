# Build stage
FROM golang:1.23 as builder

# Set the working directory inside the container
WORKDIR /app

# Copy the source code into the container
COPY . .

# Build the Go app
RUN go build ./cmd/sbom2tree

# Final stage
FROM scratch

# Copy the binary from the build stage
COPY --from=builder /app/sbom2tree /app/sbom2tree

# Set the working directory in the container
WORKDIR /app

ENTRYPOINT ["/app/sbom2tree"]
