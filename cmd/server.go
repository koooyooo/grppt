package main

import (
	"fmt"
	"log"
	"net"

	"github.com/koooyooo/grppt/serv"

	"github.com/koooyooo/grppt/pb"

	"google.golang.org/grpc"
)

func main() {
	fmt.Println("start Server")

	listener, err := net.Listen("tcp", "0.0.0.0:5051")
	if err != nil {
		log.Fatal(err)
	}
	server := grpc.NewServer()
	pb.RegisterGrpptServiceServer(server, &serv.GrpptServiceServer{})

	if err := server.Serve(listener); err != nil {
		log.Fatal(err)
	}
}
