package server

import (
	"context"
	"net/http"
	"strconv"
	"time"

	"github.com/koooyooo/grppt/core/converter"
	"github.com/koooyooo/grppt/pb"
)

type GrpptServiceServer struct{}

func (*GrpptServiceServer) Do(ctx context.Context, req *pb.Request) (*pb.Response, error) {
	return control(req)
}

func control(request *pb.Request) (*pb.Response, error) {
	st := time.Now()
	httpReq, err := converter.ConvertRequestPB2HTTP(request)
	if err != nil {
		return nil, err
	}
	httpResp, err := callBackend(httpReq)
	if err != nil {
		return nil, err
	}
	defer httpResp.Body.Close()
	resp, err := converter.ConvertResponseHTTP2PB(httpResp)
	ed := time.Now()
	resp.Headers["X-Grppt-Latency-Server"] = &pb.Values{Values: []string{strconv.Itoa(int(ed.Sub(st).Milliseconds()))}}
	return resp, err
}

func callBackend(req *http.Request) (*http.Response, error) {
	client := &http.Client{}
	st := time.Now()
	res, err := client.Do(req)
	ed := time.Now()
	res.Header.Add("X-Grppt-Latency-Server-Backend-Call", strconv.Itoa(int(ed.Sub(st).Milliseconds())))
	if err != nil {
		return nil, err
	}
	return res, nil
}
