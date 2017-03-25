# Api Gateway: api world, we neeed a good solution.

Apigateway is build by golang and go-gin. we hope find a good solution privode you.<br>


<a href="https://golang.org/">golang</a>: Who use, who know.

<a href="https://gin-gonic.github.io/gin/">go-gin</a>: It`s a good framework for golang.


<h2>How to run?</h2>

1. prepare

    <pre>
        start mongo db in your localhost, and update your mongodb info in /conf/app.conf.yml.
        You also can choise mysqldb.
    </pre>

2. startup

    <pre>
    cd workdir
    go get
    go build server.go
    go run server.go
    </pre>

If running normally, you can access<a href="http://localhost:8012">http://localhost:8012</a>

<small>Keep update to here for latest changed. Thanks for you love it.</small>

