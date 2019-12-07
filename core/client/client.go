package client

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/koooyooo/grppt/core/converter"
	"github.com/koooyooo/grppt/pb"
	"google.golang.org/grpc"
)

func CreateClient() (*grpc.ClientConn, *pb.GrpptServiceClient, error) {
	conn, err := grpc.Dial("localhost:5051", grpc.WithInsecure())
	if err != nil {
		return nil, nil, err
	}
	client := pb.NewGrpptServiceClient(conn)
	return conn, &client, nil
}

func RunClient(client pb.GrpptServiceClient, req *pb.Request) (*pb.Response, error) {
	st := time.Now()
	resp, err := client.Do(context.Background(), req)
	ed := time.Now()
	fmt.Println("Server Call", ed.Sub(st).Milliseconds()) // TODO
	return resp, err
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
