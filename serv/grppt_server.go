package serv

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"

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
	url, err := url.Parse(req.Url)
	if err != nil {
		return nil, err
	}
	reqBodyBytes := []byte(req.Body)
	httpReq := http.Request{
		Method:     req.Method,
		URL:        url,
		Proto:      "HTTP",
		ProtoMajor: 1,
		ProtoMinor: 1,
		Header:     toMap(req.Headers),
		Body:       ioutil.NopCloser(bytes.NewReader(reqBodyBytes)),
		GetBody: func() (io.ReadCloser, error) {
			return ioutil.NopCloser(bytes.NewReader(reqBodyBytes)), nil
		},
		ContentLength:    int64(len(reqBodyBytes)),
		TransferEncoding: []string{},
		Close:            false,
	}
	fmt.Println(httpReq)

	client := &http.Client{}
	res, err := client.Do(&httpReq)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	respBodyBytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	return &pb.HttpResponse{
		StatusCode:  int32(res.StatusCode),
		ReasonPhase: res.Status,
		Headers:     toValuesMap(res.Header),
		Body:        string(respBodyBytes),
	}, nil
}

func toMap(valuesMap map[string]*pb.Values) map[string][]string {
	m := map[string][]string{}
	for k, v := range valuesMap {
		m[k] = v.Values
	}
	return m
}

func toValuesMap(m map[string][]string) map[string]*pb.Values {
	valuesMap := map[string]*pb.Values{}
	for k, v := range m {
		valuesMap[k] = &pb.Values{
			Values: v,
		}
	}
	return valuesMap
}
