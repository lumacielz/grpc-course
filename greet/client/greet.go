package main

import (
	"context"
	"github.com/lumacielz/grpc-course/greet/proto"
	"log"
)

func doGreet(c proto.GreetServiceClient) {
	log.Println("doGreet was invoked")
	res, err := c.Greet(context.Background(), &proto.GreetRequest{FirstName: "Lulis"})
	if err != nil {
		log.Fatalf("Could not greet: %v\n", err)
	}

	log.Printf("Greeting: %s\n", res.Result)
}
