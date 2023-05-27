package main

import (
	"context"
	"github.com/lumacielz/grpc-course/blog/proto"
	"log"
)

func readBlog(c proto.BlogServiceClient, id string) *proto.Blog {
	log.Println("readBlog was invoked")

	req := proto.BlogId{Id: id}
	res, err := c.ReadBlog(context.Background(), &req)
	if err != nil {
		log.Printf("Error while reading: %v\n", err)
	}

	log.Println("Blog was read: %v\n", res)
	return res
}
