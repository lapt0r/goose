FROM golang:1.14.2-alpine AS build
#build source directories
RUN mkdir /app
ADD . /app
WORKDIR /app
ADD ./config .
# Add this go mod download command to pull in any dependencies
RUN go mod download
# Our project will now successfully build with the necessary go libraries included.
RUN go build -o goose .
# Our start command which kicks off
# our newly created binary executable
ENTRYPOINT []
CMD ["/app/goose"]