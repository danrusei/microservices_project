# This is a multi-stage build ("builder" create the binary, "release" run the app)
# Fist Stage: builder 
# Use base golang image from Docker Hub
FROM golang:1.12-alpine AS builder
# Install ca-certificates bundle & git inside the docker image
RUN apk add --no-cache ca-certificates git
# Sets the environment variables
#ENV PROJECT github.com/Danr17/microservices_project/section_3/frontend
# Sets the working directory
#WORKDIR /go/src/$PROJECT/
WORKDIR /microservices_project/section_3/frontend
# Copy the application source code including go.mod and go.sum files
COPY . .
# Download dependencies
RUN go mod download

# Compile the application to /frontend.
RUN go build -ldflags="-w -s" -o /frontend main.go

# Second Stage: release
FROM alpine as release
# Install ca-certificates bundle & git inside the docker image
RUN apk add --no-cache ca-certificates
# Sets the working directory
WORKDIR /frontend
# Copies new files or directories from src to dst
COPY --from=builder /frontend ./server
# Informs Docker that the container listens on the specified network ports at runtime
EXPOSE 8082
# Configure a container that will run as an executable
ENTRYPOINT ["/frontend/server"]





