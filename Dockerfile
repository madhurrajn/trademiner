# Dockerfile References: https://docs.docker.com/engine/reference/builder/

# Start from golang:1.12-alpine base image
FROM golang:1.17.3-alpine

# The latest alpine images don't have some tools like (`git` and `bash`).
# Adding git, bash and openssh to the image
RUN apk update && apk upgrade && \
    apk add --no-cache bash git openssh

# Add Maintainer Info
LABEL maintainer="Madhur Raj N <madhurrajn@gmail.com>"

# Set the Current Working Directory inside the container
WORKDIR /app
ENV GOBIN /go/bin
#RUN bash /app/src/getgetter.sh
RUN go env -w GO111MODULE=auto
COPY go.mod ./
RUN go mod download
#RUN go get github.com/tiloso/googlefinance
#RUN go get -u google.golang.org/api/sheets/v4
#RUN go get -u golang.org/x/oauth2/google
#RUN go get -u github.com/nanobox-io/golang-scribble
#RUN go get -u github.com/gin-gonic/gin
RUN unset GOPATH
RUN go mod tidy

# Copy go mod and sum files
# COPY go.mod go.sum ./

# Download all dependancies. Dependencies will be cached if the go.mod and go.sum files are not changed
# RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
COPY . /go

# Build the Go app
# RUN go build -o main .

# Expose port 8080 to the outside world
#EXPOSE 8080

# Run the executable
#CMD ["./main"]
