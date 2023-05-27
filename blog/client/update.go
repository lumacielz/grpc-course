package main

import (
	"context"
	"github.com/lumacielz/grpc-course/blog/proto"
	"log"
)

func updateBlog(c proto.BlogServiceClient, id string) {
	log.Println("updateBlog was invoked")

	newBlog := proto.Blog{
		Id:       id,
		AuthorId: "Fulano",
		Content:  "Content of the first blog with additions",
	}

	_, err := c.UpdateBlog(context.Background(), &newBlog)
	if err != nil {
		log.Fatalf("Error updating: %v\n", err)
	}
	log.Println("Blog was updated!")
}
