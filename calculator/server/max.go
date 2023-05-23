package main

import (
	"github.com/lumacielz/grpc-course/calculator/proto"
	"io"
	"log"
)

func (s *Server) Max(stream proto.CalculatorService_MaxServer) error {
	var currentMax int64
	for {
		res, err := stream.Recv()
		if err == io.EOF {
			return nil
		}

		if err != nil {
			log.Fatalf("Error receiving from client: %v\n", err)
		}

		if res.GetNum() > currentMax {
			currentMax = res.GetNum()
		}

		err = stream.Send(&proto.MaxResponse{
			Max: currentMax,
		})
		if err != nil {
			log.Fatalf("Error while sending data to client: %v\n", err)
		}
	}
	return nil
}
