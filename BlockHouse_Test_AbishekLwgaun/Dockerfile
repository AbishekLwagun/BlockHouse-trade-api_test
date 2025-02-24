# Uses the official Golang image as the base image for building the app
FROM golang:1.23-alpine as builder

# Installing necessary packages including gcc for cgo
RUN apk add --no-cache gcc musl-dev

# Setting up the working directory inside the container
WORKDIR /app

# Copy Go module files to the container
COPY go.mod go.sum ./

# Download all the dependencies
RUN go mod tidy

# Copy the rest of the application files into the container
COPY . .

# Build the Go app with CGO_ENABLED=1  #this shoud be done , it created problem first
RUN CGO_ENABLED=1 go build -o main .

# Start a new stage to create the final lightweight image
FROM alpine:latest

# Install necessary certificates to allow HTTPS requests
RUN apk --no-cache add ca-certificates

# Set the working directory for the final image
WORKDIR /root/

# Copy the compiled binary from the build stage to the final image
COPY --from=builder /app/main .

# Expose port 8080 for the app to be accessible
EXPOSE 8080

# Define the default command to run the binary when the container starts
CMD ["./main"]
