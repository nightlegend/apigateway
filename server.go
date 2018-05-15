package main

import (
	"os"

	log "github.com/Sirupsen/logrus"
	"github.com/nightlegend/apigateway/core/router"
)

// Api server start from here. router is define your api router and public it.
func main() {

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

	// Init env configure.
	// os.Setenv("APIGATEWAY_RUNING_ENV", "development")
	// go conf.InitServer()

	// Start socket server(listen on 5000).
	// go socketio.RunServer()

	router.Start() //start api server, listen on 8012.

}
