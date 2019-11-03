package converter

import (
	"bytes"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/koooyooo/grppt/pb"
)

func ConvertRequestHTTP2PB(req *http.Request) (*pb.Request, error) {
	bodyBytes, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return nil, err
	}
	return &pb.Request{
		Proto:      req.Proto,
		ProtoMajor: int32(req.ProtoMajor),
		ProtoMinor: int32(req.ProtoMinor),
		Method:     req.Method,
		Url:        req.URL.String(),
		Headers:    toValuesMap(req.Header),
		Body:       bodyBytes,
	}, nil
}

func ConvertRequestPB2HTTP(req *pb.Request) (*http.Request, error) {
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

func ConvertResponseHTTP2PB(res *http.Response) (*pb.Response, error) {
	respBodyBytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	return &pb.Response{
		Proto:       res.Proto,
		ProtoMajor:  int32(res.ProtoMajor),
		ProtoMinor:  int32(res.ProtoMajor),
		StatusCode:  int32(res.StatusCode),
		ReasonPhase: res.Status,
		Headers:     toValuesMap(res.Header),
		Body:        respBodyBytes,
	}, nil
}

func ConvertResponsePB2HTTP(res *pb.Response) (*http.Response, error) {
	return &http.Response{
		Proto:      "HTTP",
		ProtoMajor: 1,
		ProtoMinor: 1,
		StatusCode: int(res.StatusCode),
		Status:     res.ReasonPhase,
		Header:     toMap(res.Headers),
		Body:       ioutil.NopCloser(bytes.NewReader(res.Body)),
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
