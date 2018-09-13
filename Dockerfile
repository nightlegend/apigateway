FROM golang:1.10

MAINTAINER David Guo

WORKDIR /go/src/apigateway/

COPY ./ /go/src/apigateway/

RUN go get && \
    go build

EXPOSE 8080

CMD ["go", "run server.go"]