package main

import (
	"flag"
	"fmt"
	connect_four_proto "github.com/sternj/go-connect-four/internal/connect-four-proto"
	connect_four_server "github.com/sternj/go-connect-four/internal/connect-four-server"
	"google.golang.org/grpc"
	"log"
	"net"
)

var port = flag.Int("port", 8080, "port to listen on")
func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("Error listening on port %d", *port)
	}
	grpcServer := grpc.NewServer()
	connect_four_proto.RegisterConnectFourServer(grpcServer, connect_four_server.NewConnectFourService())
	fmt.Printf("Listening on port %d\n", *port)
	err2 := grpcServer.Serve(lis)
	if err2 != nil {
		log.Fatalf("error listening: %s", err2.Error())
	}
}
