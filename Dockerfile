# FROM golang:1.10
FROM grpc-go:latest

LABEL maintainer="David Guo"

WORKDIR /go/src/github.com/nightlegend/apigateway/

COPY ./ /go/src/github.com/nightlegend/apigateway/

RUN go get && \
    go build

EXPOSE 8080

CMD ["go", "run server.go"]


# docker run -it -v D:\opensource\golangspace\src\github.com\nightlegend\apigateway:/go/src/github.com/nightlegend/apigateway apigateway:gencode /bin/bash