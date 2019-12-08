package main

import (
	"fmt"
	"log"
	"net"

	"github.com/koooyooo/grppt/core/server"
	"github.com/koooyooo/grppt/pb"

	"google.golang.org/grpc"
)

func main() {
	fmt.Println("start Server")
	RunServer("tcp", "0.0.0.0:5051")
}

func RunServer(network, address string) {
	listener, err := net.Listen(network, address)
	if err != nil {
		log.Fatal(err)
	}
	grpcServ := grpc.NewServer()
	pb.RegisterGrpptServiceServer(grpcServ, &server.GrpptServiceServer{})

	if err := grpcServ.Serve(listener); err != nil {
		log.Fatal(err)
	}
}
