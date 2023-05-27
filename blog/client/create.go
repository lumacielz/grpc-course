package main

import (
	"context"
	"github.com/lumacielz/grpc-course/blog/proto"
	"log"
)

func createBlog(c proto.BlogServiceClient) string {
	log.Println("createBlog was invoked")

	blog := &proto.Blog{
		AuthorId: "Lulis",
		Title:    "My First Post",
		Content:  "some content",
	}

	res, err := c.CreateBlog(context.Background(), blog)
	if err != nil {
		log.Fatalf("Unexpected error: %v\n", err)
	}

	log.Printf("Blog has been created with Id: %s\n", res.Id)
	return res.Id
}
