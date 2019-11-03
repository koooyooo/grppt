package service

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"time"

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
	fmt.Println("???", resp)
	return resp, err
}

func callBackend(req *http.Request) (*http.Response, error) {
	client := &http.Client{}
	st := time.Now()
	res, err := client.Do(req)
	ed := time.Now()
	res.Header.Add("X-Grppt-Http-Time", strconv.Itoa(int(ed.Sub(st).Milliseconds())))
	if err != nil {
		return nil, err
	}
	return res, nil
}
