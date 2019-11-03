package service

import (
	"context"
	"fmt"
	"net/http"

	"github.com/koooyooo/grppt/converter"
	"github.com/koooyooo/grppt/pb"
)

type GrpptServiceServer struct{}

func (*GrpptServiceServer) Do(ctx context.Context, req *pb.Request) (*pb.Response, error) {
	return baseFlow(req)
}

func baseFlow(request *pb.Request) (*pb.Response, error) {
	httpReq, err := converter.ConvertRequestPB2HTTP(request)
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

func callBackend(req *http.Request) (*http.Response, error) {
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	return res, nil
}
