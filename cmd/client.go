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
		Proto:      "HTTP",
		ProtoMajor: 1,
		ProtoMinor: 1,
		Method:     "GET",
		Url:        "https://httpbin.org/get",
		Headers: map[string]*pb.Values{
			"hello": {Values: []string{"world", "baby"}},
		},
	}
	resp, err := client.Do(context.Background(), req)
	fmt.Println(resp.StatusCode)
	fmt.Println(resp.Body)
}
