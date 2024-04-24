# Builder stage
FROM golang:1.22-alpine3.19 AS builder
WORKDIR /build
# Copy the Go source code and relevant files to the /build directory
COPY . .
# Compile the application to a binary called 'main'
RUN go build -o /build/main

# Final stage
FROM ubuntu:20.04
WORKDIR /app
# Copy the compiled binary from the builder stage
COPY --from=builder /build/main .
# Copy the template file from the builder stage
COPY --from=builder /build/templates/invoice.html templates/

# Install wkhtmltopdf for converting HTML to PDF
RUN apt-get update && apt-get install -y wkhtmltopdf \
    # Clean up the apt cache to reduce the image size
    && rm -rf /var/lib/apt/lists/*

# Inform Docker that the container is listening on port 8080 at runtime
EXPOSE 8080

# Define the container's entrypoint as the application binary
ENTRYPOINT ["/app/main"]
