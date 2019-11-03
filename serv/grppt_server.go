package serv

import (
	"context"
	"fmt"
	"net/http"

	"github.com/koooyooo/grppt/converter"

	"github.com/koooyooo/grppt/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type GrpptServiceServer struct{}

func (*GrpptServiceServer) Do(ctx context.Context, req *pb.HttpRequest) (*pb.HttpResponse, error) {
	fmt.Println("receive request")
	resp, err := clientCall(req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (*GrpptServiceServer) DoStream(srv pb.GrpptService_DoStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method DoStream not implemented")
}

func clientCall(req *pb.HttpRequest) (*pb.HttpResponse, error) {
	httpReq, err := converter.ConvertRequestPB2HTTP(req)
	if err != nil {
		return nil, err
	}
	client := &http.Client{}
	res, err := client.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	return converter.ConvertResponseHTTP2PB(res)
}
