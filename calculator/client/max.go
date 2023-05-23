package main

import (
	"context"
	"github.com/lumacielz/grpc-course/calculator/proto"
	"io"
	"log"
	"time"
)

func doMax(c proto.CalculatorServiceClient) {
	stream, err := c.Max(context.Background())
	if err != nil {
		log.Fatalf("Error while creating stream: %v\n", err)
	}

	reqs := []proto.MaxRequest{
		{Num: 8},
		{Num: 7},
		{Num: 9},
	}

	waitC := make(chan struct{})
	go func() {
		for _, r := range reqs {
			log.Printf("Sending request: %d\n", r.Num)
			err := stream.Send(&r)
			if err != nil {
				log.Printf("Error sending request: %v\n", err)
			}
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
				log.Printf("Error receiving: %v\n", err)
			}
			log.Printf("Current Max: %d\n", res.Max)
		}
		close(waitC)
	}()
	<-waitC
}
