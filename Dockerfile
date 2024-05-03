# Start from the official Golang image to build our application.
FROM golang:1.22

# Copy the local package files to the container workspace.
ADD . /app

# Set the Current Working Directory inside the container
WORKDIR /app

# Build the Go app
RUN go build -o main ./cmd

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["./main"]
