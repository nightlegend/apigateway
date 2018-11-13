package grpc

import (
	"context"
	"flag"
	"log"
	"time"

	pb "github.com/nightlegend/apigateway/core/grpc/services"

	"google.golang.org/grpc"
)

var (
	serverAddrs = flag.String("server_addr", "127.0.0.1:10000", "The server address in the format of host:port")
)

// APIClient new a GRPC connection client
func APIClient() (*grpc.ClientConn, pb.RouteGuideClient, error) {
	flag.Parse()
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())

	conn, err := grpc.Dial(*serverAddrs, opts...)
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	return conn, pb.NewRouteGuideClient(conn), err
}

// Testing node deregister is delete service from consul
func Testing(client pb.RouteGuideClient, r *pb.Request) *pb.Response {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	resp, err := client.Testing(ctx, r)
	if err != nil {
		log.Fatalf("%v.Testing(_) = _, %v: ", client, err)
	}
	return resp
}
