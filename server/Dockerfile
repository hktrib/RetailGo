# Use the official Golang image as the base image
FROM golang:latest

# Set the working directory inside the container
WORKDIR /app

# Copy the entire project into the container
COPY . .

# Build the Go application
RUN go build -o main ./cmd/

# Expose the port that your web application will run on
EXPOSE 8080

# Command to run the executable
CMD ["./main"]
