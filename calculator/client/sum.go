package main

import (
	"context"
	"github.com/lumacielz/grpc-course/calculator/proto"
	"log"
)

func doSum(c proto.CalculatorServiceClient) {
	var res, err = c.Sum(context.Background(), &proto.SumRequest{Num1: 3, Num2: 11})
	if err != nil {
		log.Fatalf("Could not calculator: %v\n", err)
	}
	log.Printf("The calculator is: %d", res.Sum)
}
