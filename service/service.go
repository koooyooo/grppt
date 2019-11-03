package service

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

func (*GrpptServiceServer) Do(ctx context.Context, req *pb.Request) (*pb.Response, error) {
	httpReq, err := converter.ConvertRequestPB2HTTP(req)
	fmt.Println(httpReq)
	if err != nil {
		return nil, err
	}
	httpResp, err := callBackend(httpReq)
	if err != nil {
		return nil, err
	}
	defer httpResp.Body.Close()
	resp, err := converter.ConvertResponseHTTP2PB(httpResp)
	return resp, err
}

func (*GrpptServiceServer) DoStream(srv pb.GrpptService_DoStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method DoStream not implemented")
}

func callBackend(req *http.Request) (*http.Response, error) {
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	return res, nil
}
