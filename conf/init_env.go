package conf

import (
	"io/ioutil"
	"os"

	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
)

// InitServer is init server configure.
func InitServer() {
	/*
	 * Init global logs file
	 */
	// execDirAbsPath, _ := os.Getwd()
	// f, err := os.OpenFile(execDirAbsPath+"/logs/app.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// //defer to close when you're done with it, not because you think it's idiomatic!
	// defer f.Close()
	// //set output of logs to f
	// log.SetOutput(f)
	execDirAbsPath, _ := os.Getwd()
	log.Info("start init env configure")
	env := os.Getenv("APIGATEWAY_RUNING_ENV")
	log.Info("You load env is:" + env)

	data, err := ioutil.ReadFile(execDirAbsPath + "/conf/env/" + env + ".conf.yaml")
	if err != nil {
		log.Println(err)
	}
	var config *Config
	err = yaml.Unmarshal([]byte(data), &config)
	if err != nil {
		log.Info(err)
	}
	log.Info(config.Mongohost)
}
