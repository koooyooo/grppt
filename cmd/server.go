package main

import (
	"fmt"
	"log"
	"net"

	"github.com/koooyooo/grppt/pb"
	"github.com/koooyooo/grppt/service"

	"google.golang.org/grpc"
)

func main() {
	fmt.Println("start Server")
	RunServer()
}

func RunServer() {
	listener, err := net.Listen("tcp", "0.0.0.0:5051")
	if err != nil {
		log.Fatal(err)
	}
	grpcServ := grpc.NewServer()
	pb.RegisterGrpptServiceServer(grpcServ, &service.GrpptServiceServer{})

	if err := grpcServ.Serve(listener); err != nil {
		log.Fatal(err)
	}
}
