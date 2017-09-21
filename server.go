package main

import (
	"apigateway/conf"
	"apigateway/core/router"
	"apigateway/core/socketio"
	"log"
	"os"
)

/*
 * Author: David Guo
 * Start Application
 * * * 1. Start RESTFUL
 * * * 2. Start Socket
 */
func main() {
	/*
	 * Init global logs file
	 */
	execDirAbsPath, _ := os.Getwd()
	f, err := os.OpenFile(execDirAbsPath+"/logs/app.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Fatal(err)
	}
	//defer to close when you're done with it, not because you think it's idiomatic!
	defer f.Close()
	//set output of logs to file
	log.SetOutput(f)

	/*
	 * Init env configure.
	 */
	os.Setenv("APIGATEWAY_RUNING_ENV", "development")
	go conf.InitServer()

	/*
	 * Init RESTFul server
	 * Init Socket server
	 */

	// Start socket server(listen on 5000).
	go socketio.RunServer()
	// Start RESTFul server (listen on 8012)
	router.Start()

}
