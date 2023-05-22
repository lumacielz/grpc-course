package main

import (
	"fmt"
	"github.com/lumacielz/grpc-course/greet/proto"
	"io"
	"log"
)

func (s Server) LongGreet(stream proto.GreetService_LongGreetServer) error {
	log.Println("LongGreet function was invoked")
	res := ""
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(
				&proto.GreetResponse{
					Result: res,
				})
		}

		if err != nil {
			log.Fatalf("Error reading client stream: %v\n", err)
		}

		res += fmt.Sprintf("Hello %s!\n", req.FirstName)
	}
}
