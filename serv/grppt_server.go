package serv

import (
	"context"
	"fmt"

	"github.com/koooyooo/grppt/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type GrpptServiceServer struct{}

func (*GrpptServiceServer) Do(ctx context.Context, req *pb.HttpRequest) (*pb.HttpResponse, error) {
	fmt.Println("receive request")
	return &pb.HttpResponse{
		Version: "1.1",
	}, nil
}

func (*GrpptServiceServer) DoStream(srv pb.GrpptService_DoStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method DoStream not implemented")
}
