package client

import (
	"context"
	"net/http"

	"github.com/koooyooo/grppt/core/converter"
	"github.com/koooyooo/grppt/pb"
	"google.golang.org/grpc"
)

func CreateClient(target string) (*grpc.ClientConn, *pb.GrpptServiceClient, error) {
	conn, err := grpc.Dial(target, grpc.WithInsecure())
	if err != nil {
		return nil, nil, err
	}
	client := pb.NewGrpptServiceClient(conn)
	return conn, &client, nil
}

func RunClient(client pb.GrpptServiceClient, req *pb.Request) (*pb.Response, error) {
	return client.Do(context.Background(), req)
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
