package main

import (
	"flag"
	"fmt"
	"net"
	"os"

	mgrpc "github.com/nightlegend/apigateway/core/grpc"

	log "github.com/Sirupsen/logrus"
	pb "github.com/nightlegend/apigateway/core/grpc/services"
	"github.com/nightlegend/apigateway/core/router"
	"google.golang.org/grpc"
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

	// GRPC
	//
	// Here will enable grpc server, if you don`t want it, you can disable it
	go func() {
		lis, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%d", 10000))
		if err != nil {
			log.Fatalf("failed to listen: %v", err)
		}
		var opts []grpc.ServerOption
		grpcServer := grpc.NewServer(opts...)
		pb.RegisterRouteGuideServer(grpcServer, mgrpc.NewServer())
		grpcServer.Serve(lis)
	}()

	// HTPP
	//
	// start api server, *env is what`s environment will running, currentlly this only for enable or disable debug modle
	// After may be use it load different varible.
	router.Start(*env)
}
