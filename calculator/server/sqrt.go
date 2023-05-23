package main

import (
	"context"
	"fmt"
	"github.com/lumacielz/grpc-course/calculator/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"math"
)

func (s *Server) Sqrt(ctx context.Context, req *proto.SqrtRequest) (*proto.SqrtResponse, error) {
	log.Printf("Sqrt was invoked with: %v\n", req)
	num := req.Num
	if num < 0 {
		return nil, status.Errorf(
			codes.InvalidArgument,
			fmt.Sprintf("Received a negative number"),
		)
	}
	return &proto.SqrtResponse{
		Res: math.Sqrt(float64(req.GetNum())),
	}, nil
}
