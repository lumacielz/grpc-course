package main

import (
	"github.com/lumacielz/grpc-course/calculator/proto"
	"io"
	"log"
)

func (s *Server) Average(stream proto.CalculatorService_AverageServer) error {
	log.Println("Average function was invoked")
	var sum float64
	var counter int
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(
				&proto.AverageResponse{
					Res: sum / float64(counter),
				})
		}

		if err != nil {
			log.Fatalf("Error reading client stream: %v\n", err)
		}

		sum += float64(req.Num)
		counter++
	}
	return nil
}
