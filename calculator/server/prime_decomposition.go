package main

import (
	"github.com/lumacielz/grpc-course/calculator/proto"
	"log"
)

func (s *Server) PrimeDecomposition(in *proto.PrimeRequest, stream proto.CalculatorService_PrimeDecompositionServer) error {
	num := in.GetNum()
	log.Println("PrimeDecomposition function called with %d", num)
	var k int64 = 2
	for num > 1 {
		if num%k == 0 {
			stream.Send(&proto.PrimeResponse{Prime: k})
			num = num / k
		} else {
			k += 1
		}
	}

	return nil
}
