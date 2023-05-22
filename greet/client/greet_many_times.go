package main

import (
	"context"
	"github.com/lumacielz/grpc-course/greet/proto"
	"io"
	"log"
)

func doGreetManyTimes(c proto.GreetServiceClient) {
	log.Println("doGreetManyTimes was invoked")

	req := &proto.GreetRequest{
		FirstName: "Lulis",
	}

	stream, err := c.GreetManyTimes(context.Background(), req)
	if err != nil {
		log.Fatalf("Error calling GreetManyTimes: %v\n", err)
	}

	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("Error while reading the stream: %v\n", err)
		}

		log.Printf("GreetManyTimes: %s\n", msg.Result)
	}
}
