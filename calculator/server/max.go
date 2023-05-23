package main

import (
	"github.com/lumacielz/grpc-course/calculator/proto"
	"io"
	"log"
	"math"
)

func (s *Server) Max(stream proto.CalculatorService_MaxServer) error {
	var currentMax float64
	for {
		res, err := stream.Recv()
		if err == io.EOF {
			return nil
		}

		if err != nil {
			log.Fatalf("Error receiving from client: %v\n", err)
		}

		currentMax = math.Max(float64(currentMax), float64(res.Num))
		err = stream.Send(&proto.MaxResponse{
			Max: int64(currentMax),
		})
		if err != nil {
			log.Fatalf("Error while sending data to client: %v\n", err)
		}
	}
	return nil
}
