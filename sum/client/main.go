package main

import (
	"context"
	"github.com/lumacielz/grpc-course/sum/proto"
	"google.golang.org/grpc"
	"log"
)

var addr string = "localhost:50051"

func main() {
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect: %v\n", err)
	}
	defer conn.Close()

	c := proto.NewSumServiceClient(conn)
	res, err := c.Sum(context.Background(), &proto.SumRequest{Num1: 3, Num2: 11})
	if err != nil {
		log.Fatalf("Could not sum: %v\n", err)
	}
	log.Printf("The sum is: %d", res.Sum)
}
