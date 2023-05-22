package main

import (
	"context"
	"github.com/lumacielz/grpc-course/calculator/proto"
	"log"
)

func doAverage(c proto.CalculatorServiceClient) {
	numbers := []proto.AverageRequest{
		{Num: 7},
		{Num: 9},
		{Num: 8},
		{Num: 4},
		{Num: 6},
	}
	stream, err := c.Average(context.Background())
	if err != nil {
		log.Fatalf("Error calling Average: %v\n", err)
	}

	for _, n := range numbers {
		log.Printf("Sending %d...\n", n.Num)
		stream.Send(&n)
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Error while receiving response from Average: %v\n", err)
	}

	log.Printf("Average: %.2f\n", res.Res)
}
