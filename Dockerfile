# Use the official Go image as the base image
FROM golang:1.21.0 AS build

# Set the working directory inside the container
WORKDIR /app

# Copy the Go application source code to the container
COPY . .

# Build the Go application
RUN CGO_ENABLED=0 GOOS=linux go build -o app

# Use a smaller base image for the final image
FROM alpine:latest

# Set the working directory inside the container
WORKDIR /app

# Copy the built executable from the previous build stage
COPY --from=build /app/app .

# Run the Go application
CMD ["./app"]
