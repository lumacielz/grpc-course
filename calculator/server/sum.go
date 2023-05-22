package main

import (
	"context"
	"github.com/lumacielz/grpc-course/calculator/proto"
	"log"
)

func (s Server) Sum(ctx context.Context, req *proto.SumRequest) (*proto.SumResponse, error) {
	log.Printf("Sum function was invoked with params %v and %v", req.Num1, req.Num2)
	return &proto.SumResponse{Sum: req.GetNum1() + req.GetNum2()}, nil
}
