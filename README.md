# Api Gateway: Try a new way by golang

[![Build Status](https://travis-ci.org/nightlegend/apigateway.svg?branch=dev1.0)](https://travis-ci.org/nightlegend/apigateway) [![codecov](https://codecov.io/gh/nightlegend/apigateway/branch/dev1.0/graph/badge.svg)](https://codecov.io/gh/nightlegend/apigateway) [![Go Report Card](https://goreportcard.com/badge/github.com/nightlegend/apigateway)](https://goreportcard.com/report/github.com/nightlegend/apigateway) [![Gitter](https://badges.gitter.im/nightlegend/apigateway.svg)](https://gitter.im/nightlegend/apigateway?utm_source=badge&utm_medium=badge&utm_campaign=pr-badge)

Apigateway is a api-gateway server demo written in [golang](https://golang.org/) and [go-gin](https://gin-gonic.github.io/gin/). It features a simple and better performance, you can faster build your api server by this templete. If you need fast build a api gateway server, you will love Apigateway.


<h1>Design</h1>

![Gopher image](doc/structure.jpg)

<h2>Define</h2>

<h5>Router: control all http request, and dispatch each requets to services.
<h5>Api: Handle all request and provide a service to router.
<h5>Worker: Here is define some backand task/job do somthing.
<h5>Socket: Provide a socket server.

<h1>How to run ?</h1>

<h2>Prepare step(optional)</h2>

>start mongo db in your localhost, and update your mongodb info in /conf/app.conf.yml.You also can select mysql.


<h2>Start APIGATEWAY</h2>

* Init workdir
```sh
git clone https://github.com/nightlegend/apigateway.git
go get github.com/kardianos/govendor
cd $GOPATH/src/github.com/nightlegend/apigateway
govendor init
govendor add +external
govendor install +local
```
> if can`t recognize govendor, please try $GOPATH/bin/govendor.

* Start APIGATEWAY
```sh
# start with default
go run server.go
# -env: current for enable/disable debug model.
go run server.go -env development
```


If running normally, you can access<a href="http://localhost:8080">http://localhost:8080</a>

**Application details**

---

1. Server starting from server.go
   ```go
    package main
    import (
        "flag"
        "os"
        log "github.com/Sirupsen/logrus"
        "github.com/nightlegend/apigateway/core/router"
    )

    var (
        env = flag.String("env", "development", "running environment")
    )

    // Api server start from here. router is define your api router and public it.
    func main() {
        flag.Parse()
        // set golable logs file path.
        execDirAbsPath, _ := os.Getwd()
        f, err := os.OpenFile(execDirAbsPath+"/logs/app.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
        if err != nil {
            log.Fatal(err)
        }
        //defer to close when you're done with it, not because you think it's idiomatic!
        defer f.Close()
        //set output of logs to file
        log.SetOutput(f)

        // start api server, *env is what`s environment will running, currentlly this only for enable or disable debug modle
        // After may be use it load different varible.
        router.Start(*env)
    }
   ```
2. Project code structure

    | folder        | content                                   |
    | ------------- |:-------------                             |
    | conf          | put some application configure to here    |
    | core          | put core sources to here                  |
    | middleware    | put middleware code to here, like cors    |
    | logs          | here will save console log                |
    | vendor        | here is save third party                  |
    | doc           | here is save some document about project  |

3. Router define
   
   Router, It`s like your application gate, help you dispatch all request to target service.
   >go to apigateway/core/router/router.go, you can define your router.

    ```go
    func Start(env string) {
        // running mode switcher
        switch env {
        case "development":
            gin.SetMode(gin.DebugMode)
        default:
            gin.SetMode(gin.ReleaseMode)
        }
        router := gin.New()
        router.Use(middleware.CORSMiddleware())
        router.Use(gin.Logger())
        //No Permission Validation
        public.APIRouter(router)
        //Permission Validation
        private.APIRouter(router)
        router.Run(LisAddr)
    }
    ```
    
**Related project** 

---

If you need a [frond-end](https://github.com/nightlegend/Dashboard) template, It`s will be help you.

If you need a [SOCKET-SERVER](https://github.com/nightlegend/hi), It`s can help you.

If you interest grpc, I am happy to give a sample to you, [GRPC-GO](https://github.com/nightlegend/grpc-server-go), hope you love it, thanks.

