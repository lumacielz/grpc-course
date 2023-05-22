package main

import (
	"github.com/lumacielz/grpc-course/greet/proto"
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

	c := proto.NewGreetServiceClient(conn)
	doGreet(c)
}
