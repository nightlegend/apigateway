FROM golang:1.7

MAINTAINER david.guo18@yahoo.com

COPY ./ /go/src/magic/

WORKDIR /go/src/magic/

RUN go get && \
    go build

EXPOSE 8089

CMD ["go run", "main.go"]