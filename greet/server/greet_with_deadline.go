package main

import (
	"context"
	"github.com/lumacielz/grpc-course/greet/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"time"
)

func (s *Server) GreetWithDeadline(ctx context.Context, in *proto.GreetRequest) (*proto.GreetResponse, error) {
	log.Printf("GreetWithDeadline was invoked with: %v\n", in)

	for i := 0; i < 3; i++ {
		if ctx.Err() == context.DeadlineExceeded {
			log.Println("The client cancelled the request")
			return nil, status.Error(codes.DeadlineExceeded, "The cliend cancelled the request")
		}
		time.Sleep(1 * time.Second)
	}

	return &proto.GreetResponse{
		Result: "Hello " + in.FirstName,
	}, nil
}
