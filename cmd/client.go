package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/koooyooo/grppt/converter"

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
	req := &pb.Request{
		Proto:      "HTTP",
		ProtoMajor: 1,
		ProtoMinor: 1,
		Method:     "GET",
		Url:        "https://httpbin.org/get",
		Headers: map[string]*pb.Values{
			"hello": {Values: []string{"world", "baby"}},
		},
		Body: []byte("Hello Body"),
	}
	resp, err := RunClient(client, req)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(resp.StatusCode)
	fmt.Println(string(resp.Body))
}

func RunHttpClient(client pb.GrpptServiceClient, req *http.Request) (*http.Response, error) {
	pbReq, err := converter.ConvertRequestHTTP2PB(req)
	if err != nil {
		return nil, err
	}
	pbResp, err := RunClient(client, pbReq)
	if err != nil {
		return nil, err
	}
	return converter.ConvertResponsePB2HTTP(pbResp)
}

func RunClient(client pb.GrpptServiceClient, req *pb.Request) (*pb.Response, error) {
	return client.Do(context.Background(), req)
}

// FIXME: Need to assign right Request / Response set using Hash.
func RunStreamClient(client pb.GrpptServiceClient, req *pb.Request) (*pb.Response, error) {
	stream, err := client.DoStream(context.Background())
	if err != nil {
		return nil, err
	}
	err = stream.Send(req)
	if err != nil {
		return nil, err
	}
	return stream.Recv()
}
