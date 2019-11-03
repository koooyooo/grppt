package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/koooyooo/grppt/converter"

	"github.com/koooyooo/grppt/pb"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Run Client")
	st := time.Now()
	conn, client, err := CreateClient()
	md := time.Now()
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
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
	resp, err := RunClient(*client, req)
	if err != nil {
		log.Fatal(err)
	}
	ed := time.Now()
	fmt.Println(resp.StatusCode)
	fmt.Println(resp.Headers)
	fmt.Println(string(resp.Body))
	fmt.Println("total", ed.Sub(st).Milliseconds(), "dialing", md.Sub(st).Milliseconds())
}

func CreateClient() (*grpc.ClientConn, *pb.GrpptServiceClient, error) {
	conn, err := grpc.Dial("localhost:5051", grpc.WithInsecure())
	if err != nil {
		return nil, nil, err
	}
	client := pb.NewGrpptServiceClient(conn)
	return conn, &client, nil
}

func RunHttpClient(client *pb.GrpptServiceClient, req *http.Request) (*http.Response, error) {
	pbReq, err := converter.ConvertRequestHTTP2PB(req)
	if err != nil {
		return nil, err
	}
	pbResp, err := RunClient(*client, pbReq)
	if err != nil {
		return nil, err
	}
	return converter.ConvertResponsePB2HTTP(pbResp)
}

func RunClient(client pb.GrpptServiceClient, req *pb.Request) (*pb.Response, error) {
	return client.Do(context.Background(), req)
}
