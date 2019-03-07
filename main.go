package main

import (
	"flag"
	"fmt"
	"net"    
    pb "github.com/Syati/pbsample/protos"
	"github.com/Syati/pbsample/server"
	"google.golang.org/grpc"    
	"log"
)

var (
	port = flag.Int("port", 10000, "The server port")
)

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	// enable tls ref https://github.com/grpc/grpc-go/blob/master/examples/route_guide/server/server.go
	grpcServer := grpc.NewServer()
  	newServer := &server.UserApiServer{}
	pb.RegisterUserApiServer(grpcServer, newServer)    
	grpcServer.Serve(lis)
}
