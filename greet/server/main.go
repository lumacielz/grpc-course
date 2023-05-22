package main

import (
	"github.com/lumacielz/grpc-course/greet/proto"
	"google.golang.org/grpc"
	"log"
	"net"
)

var addr string = "0.0.0.0:50051"

type Server struct {
	proto.GreetServiceServer
}

func main() {
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("Failed tp listen on %v\n", err)
	}

	log.Printf("Listening on %s\n", addr)

	s := grpc.NewServer()
	proto.RegisterGreetServiceServer(s, &Server{})

	if err = s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
