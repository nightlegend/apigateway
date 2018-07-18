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
