package main

import (
	"context"
	"github.com/lumacielz/grpc-course/greet/proto"
	"log"
	"time"
)

func doLongGreet(c proto.GreetServiceClient) {
	log.Println("doLongGreet was invoked")
	reqs := []*proto.GreetRequest{
		{FirstName: "Fulana 1"},
		{FirstName: "Fulana 2"},
		{FirstName: "Fulana 3"},
	}
	stream, err := c.LongGreet(context.Background())
	if err != nil {
		log.Fatalf("Error while calling LongGreet %v\n", err)
	}

	for _, req := range reqs {
		log.Printf("Sending req: %v\n", req)
		stream.Send(req)
		time.Sleep(1 * time.Second)
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Error while receiving response from LongGreet: %v\n", err)
	}

	log.Printf("LongGreet: %s\n", res.Result)
}
