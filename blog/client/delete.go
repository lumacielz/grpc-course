package main

import (
	"context"
	"github.com/lumacielz/grpc-course/blog/proto"
	"log"
)

func deleteBlog(c proto.BlogServiceClient, id string) {
	log.Println("deleteBlog was invoked")
	_, err := c.DeleteBlog(context.Background(), &proto.BlogId{Id: id})

	if err != nil {
		log.Fatalf("Error happened while deleting: %v\n", err)
	}

	log.Println("Blog was deleted")
}
