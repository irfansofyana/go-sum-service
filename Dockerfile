# Base image
FROM golang AS builder

# Set working directory to /src
WORKDIR /src

# Copy all files into /src
COPY . /src

# Install all dependencies
RUN go mod download

# Build application
RUN go build -o /app main.go

# Create deployment image from the built result
FROM scratch
COPY --from=builder /app .
EXPOSE 8080
CMD ["/app"]