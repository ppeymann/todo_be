# Use the official Golang image as a build environment
FROM golang:alpine as build-env

# Set the current working directory inside the container
WORKDIR /app

# Copy the source code to the container
COPY . ./

# Download the Go module dependencies
RUN go mod download

# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o /todo ./cmd/todo/main.go

# Use the official Alpine image as the base image for the final stage
FROM alpine:latest

# Set the current working directory inside the container
WORKDIR /

# Create necessary directories
RUN mkdir /data

# Add user and group
RUN addgroup --system todo && adduser -S -s /bin/false -G todo todo

# Copy the built binary from the build stage
COPY --from=build-env /todo /todo

# Copy the config folder to the root directory of the container
COPY ./config/config.json /config/config.json

COPY ./schemas /schemas

# Change ownership to the todo user
RUN chown -R todo:todo /todo
RUN chown -R todo:todo /data
RUN chown -R todo:todo /config
RUN chown -R todo:todo /schemas

# Switch to the todo user
USER todo

# Expose the necessary port
EXPOSE 8080

# Set the entrypoint to the binary
ENTRYPOINT [ "/todo" ]
