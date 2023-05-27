package main

import (
	"context"
	"github.com/lumacielz/grpc-course/blog/proto"
	"google.golang.org/protobuf/types/known/emptypb"
	"io"
	"log"
)

func listBlog(c proto.BlogServiceClient) {
	log.Println("listBlog was invoked")

	stream, err := c.ListBlogs(context.Background(), &emptypb.Empty{})
	if err != nil {
		log.Fatalf("Error while calling ListBlogs: %v\n", err)
	}

	for {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("Something Happened: %v\n", err)
		}

		log.Println(res)
	}
}
