# Use the official Golang image as the base image
FROM golang:latest

# Set the working directory inside the container
WORKDIR /app

# Copy the source code into the container
COPY . .

# Build the executable
RUN go build -o load-generator .

# Expose port 8080 (not really used in this case, but just to match your other example)
EXPOSE 8080

# Run the executable when the container starts
CMD ["./load-generator"]
