FROM golang:1.14.2-alpine AS build
#build source directories
RUN mkdir /app
ADD . /app
WORKDIR /app
ADD ./config .
# pull in Go modules and
RUN go mod download; go build -o goose .
# Our start command which kicks off
# our newly created binary executable
ENTRYPOINT []
CMD ["/app/goose"]