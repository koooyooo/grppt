package main

import (
	"context"
	"fmt"
	"log"

	"github.com/koooyooo/grppt/pb"

	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Run Client")
	conn, err := grpc.Dial("localhost:5051", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client := pb.NewGrpptServiceClient(conn)
	req := &pb.HttpRequest{
		Version: "1.1",
	}
	client.Do(context.Background(), req)
}
