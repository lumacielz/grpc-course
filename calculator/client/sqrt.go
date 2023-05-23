package main

import (
	"context"
	"github.com/lumacielz/grpc-course/calculator/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
)

func doSqrt(c proto.CalculatorServiceClient, num int64) {
	log.Println("doSqrt was invoked")
	res, err := c.Sqrt(context.Background(), &proto.SqrtRequest{Num: num})
	if err != nil {
		//Ok checks if err is a grpc error
		e, ok := status.FromError(err)
		if ok {
			log.Printf("Error message from server: %s\n", e.Message())
			if e.Code() == codes.InvalidArgument {
				log.Println("You probably sent a negative number...")
				return
			}
		} else {
			log.Fatalf("A non gRPC error: %v\n", err)
		}
	}

	log.Printf("Sqrt: %f\n", res.Res)
}
