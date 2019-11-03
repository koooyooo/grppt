package converter

import (
	"bytes"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/koooyooo/grppt/pb"
)

func ConvertRequestPB2HTTP(req *pb.HttpRequest) (*http.Request, error) {
	url, err := url.Parse(req.Url)
	if err != nil {
		return nil, err
	}
	reqBodyBytes := []byte(req.Body)
	reqBodyStream := ioutil.NopCloser(bytes.NewReader(reqBodyBytes))
	return &http.Request{
		Proto:      req.Proto,
		ProtoMajor: int(req.ProtoMajor),
		ProtoMinor: int(req.ProtoMinor),
		Method:     req.Method,
		URL:        url,
		Header:     toMap(req.Headers),
		Body:       reqBodyStream,
		GetBody: func() (io.ReadCloser, error) {
			return reqBodyStream, nil
		},
		ContentLength:    int64(len(reqBodyBytes)),
		TransferEncoding: []string{},
		Close:            false,
	}, nil
}

func ConvertResponseHTTP2PB(res *http.Response) (*pb.HttpResponse, error) {
	respBodyBytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	return &pb.HttpResponse{
		Proto:       "HTTP",
		ProtoMajor:  1,
		ProtoMinor:  1,
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