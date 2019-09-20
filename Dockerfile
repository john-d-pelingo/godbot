# Use the golang@1.13.0 base image
FROM golang:1.13.0-alpine3.10 AS base
LABEL maintainer="John Darryl Pelingo <johndarrylpelingo@gmail.com> (https://github.com/john-d-pelingo)"

FROM base as linux-packages
# Update Alpine Linux package management and add bash git and openssh.
RUN apk update && apk upgrade && apk add --no-cache bash git openssh

FROM linux-packages as go-modules
# Create a /app directory.
RUN mkdir /app
ADD go.mod /app/
ADD go.sum /app/
# Execute further commands inside the /app directory.
WORKDIR /app
RUN go mod download

FROM go-modules as app-builder
# Copy everything in the root directory into the /app directory.
ADD . /app/
# Execute further commands inside the /app directory.
WORKDIR /app
# Compile the binary executable Go program.
RUN go build -o main .

FROM app-builder as run
# Start the executable command.
CMD ["/app/main"]
