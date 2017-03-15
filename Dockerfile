FROM golang:1.7

MAINTAINER david.guo18@yahoo.com

COPY ./ /go/src/apigateway/

WORKDIR /go/src/apigateway/

RUN go get && \
    go build

EXPOSE 8089

CMD ["go run", "main.go"]