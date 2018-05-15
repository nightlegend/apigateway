# Api Gateway: Try a new way by golang

[![Build Status](https://travis-ci.org/nightlegend/apigateway.svg?branch=dev1.0)](https://travis-ci.org/nightlegend/apigateway) [![codecov](https://codecov.io/gh/nightlegend/apigateway/branch/dev1.0/graph/badge.svg)](https://codecov.io/gh/nightlegend/apigateway) [![Gitter](https://badges.gitter.im/nightlegend/apigateway.svg)](https://gitter.im/nightlegend/apigateway?utm_source=badge&utm_medium=badge&utm_campaign=pr-badge)

Apigateway is a api-gateway server demo written in [golang](https://golang.org/) and [go-gin](https://gin-gonic.github.io/gin/). It features a simple and better performance, you can faster build your api server by this templete. If you need fast build a api gateway server, you will love Apigateway.

https://travis-ci.org/nightlegend/apigateway.svg?branch=dev1.0

<h1>Design</h1>

![Gopher image](doc/structure.jpg)

<h2>Define</h2>

<h5>Router: control all http request, and dispatch each requets to  api.
<h5>Api: Handle all request and provide a service to Router.
<h5>Worker: Define some task for handle request task.
<h5>Socket: Provide a socket server.

<h1>How to running?</h1>

1. Prepare step(optional)

    <pre>
    start mongo db in your localhost, and update your mongodb info in /conf/app.conf.yml.
    You also can choise mysqldb.
    </pre>

2. Build and run step

    <pre>
    git clone https://github.com/nightlegend/apigateway.git ${golang_workspace}/src/github.com/nightlegend/
    cd ${golang_workspace}/src/github.com/nightlegend/apigateway
    go get
    go install
    go run server.go
    </pre>

If running normally, you can access<a href="http://localhost:8012">http://localhost:8012</a>


If you need a [frond-end](https://github.com/nightlegend/Dashboard) template, It`s will be help you.

<small>Keep update to here for latest changed. Thanks for you love it.</small>

