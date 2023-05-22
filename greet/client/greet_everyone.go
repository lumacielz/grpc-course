package main

import (
	"context"
	"github.com/lumacielz/grpc-course/greet/proto"
	"io"
	"log"
	"time"
)

func doGreetEveryone(c proto.GreetServiceClient) {
	log.Println("doGreetEveryone was invoked")
	stream, err := c.GreetEveryone(context.Background())
	if err != nil {
		log.Fatalf("Error while creating stream: %v\n", err)
	}

	reqs := []*proto.GreetRequest{
		{FirstName: "Fulana 1"},
		{FirstName: "Fulana 2"},
		{FirstName: "Fulana 3"},
	}

	waitC := make(chan struct{})
	go func() {
		for _, r := range reqs {
			log.Printf("Send request: %v\n", r)
			stream.Send(r)
			time.Sleep(1 * time.Second)
		}
		stream.CloseSend()
	}()

	go func() {
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				break
			}

			if err != nil {
				log.Printf("Error while receiving: %s", err)
			}

			log.Printf("Received: %v\n", res.Result)
		}
		close(waitC)
	}()

	<-waitC
}
