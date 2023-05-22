package main

import (
	"context"
	"github.com/lumacielz/grpc-course/calculator/proto"
	"io"
	"log"
)

func decomposeIntoPrimes(c proto.CalculatorServiceClient) {
	stream, err := c.PrimeDecomposition(context.Background(), &proto.PrimeRequest{Num: 2936})
	if err != nil {
		log.Fatalf("Could not decompose: %v\n", err)
	}

	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("Error while reading the stream: %v\n", err)
		}

		log.Printf("%d\n", msg.Prime)
	}
}
