# Api Gateway: Try a new way by golang

[![Build Status](https://travis-ci.org/nightlegend/apigateway.svg?branch=dev1.0)](https://travis-ci.org/nightlegend/apigateway) [![codecov](https://codecov.io/gh/nightlegend/apigateway/branch/dev1.0/graph/badge.svg)](https://codecov.io/gh/nightlegend/apigateway) [![Go Report Card](https://goreportcard.com/badge/github.com/nightlegend/apigateway)](https://goreportcard.com/report/github.com/nightlegend/apigateway) [![Gitter](https://badges.gitter.im/nightlegend/apigateway.svg)](https://gitter.im/nightlegend/apigateway?utm_source=badge&utm_medium=badge&utm_campaign=pr-badge)

Apigateway is a api-gateway server demo written in [golang](https://golang.org/) and [go-gin](https://gin-gonic.github.io/gin/). It features a simple and better performance, you can faster build your api server by this templete. If you need fast build a api gateway server, you will love Apigateway.


<h1>Design</h1>

![Gopher image](doc/structure.jpg)

<h2>Define</h2>

<h5>Router: control all http request, and dispatch each requets to  api.
<h5>Api: Handle all request and provide a service to Router.
<h5>Worker: Define some task for handle request task.
<h5>Socket: Provide a socket server.

<h1>How to running?</h1>

<h2>Prepare step(optional)</h2>
<pre>
start mongo db in your localhost, and update your mongodb info in /conf/app.conf.yml.
You also can choise mysqldb.
</pre>

<h2>Running APIGATEWAY</h2>

* Init workdir
```sh
go get github.com/kardianos/govendor
cd $GOPATH/src/github.com/nightlegend/apigateway
govendor init
govendor add +external
govendor install +local
```
> if can`t recognize govendor, please try $GOPATH/bin/govendor.

* Start GRPC-SERVER-GO
```sh
# -env mean you will up which environment.
go run server.go -env development
```


If running normally, you can access<a href="http://localhost:8080">http://localhost:8080</a>


If you need a [frond-end](https://github.com/nightlegend/Dashboard) template, It`s will be help you.

If you need a [SOCKET-SERVER](https://github.com/nightlegend/hi), It`s can help you.

If you interest grpc, I am happy to give a sample to you, [GRPC-GO](https://github.com/nightlegend/grpc-server-go), hope you love it, thanks. (development stage)

<small>Keep update to here for latest changed. Thanks for you love it.</small>

