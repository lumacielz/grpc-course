package main

import (
	"github.com/lumacielz/grpc-course/greet/proto"
	"io"
	"log"
)

func (s *Server) GreetEveryone(stream proto.GreetService_GreetEveryoneServer) error {
	log.Println("greetEveryone was invoked")

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}

		if err != nil {
			log.Fatalf("Error while reading client stream: %v\n", err)
		}

		res := "Hello " + req.FirstName + "!"
		err = stream.Send(&proto.GreetResponse{
			Result: res,
		})
		if err != nil {
			log.Fatalf("Error while sending data to client: %v\n", err)
		}
	}
}
