package main

import (
	"github.com/lumacielz/grpc-course/calculator/proto"
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

	c := proto.NewCalculatorServiceClient(conn)
	//decomposeIntoPrimes(c)
	doMax(c)
}
