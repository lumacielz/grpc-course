package main

import (
	"context"
	"fmt"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/lumacielz/grpc-course/blog/proto"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
)

func (s *Server) ListBlogs(in *empty.Empty, stream proto.BlogService_ListBlogsServer) error {
	log.Println("ListBlogs was invoked")

	cursor, err := collection.Find(context.Background(), primitive.D{{}})
	if err != nil {
		return status.Errorf(
			codes.Internal,
			fmt.Sprintf("Unknown internal error: %v\n", err),
		)
	}
	defer cursor.Close(context.Background())
	for cursor.Next(context.Background()) {
		data := &BlogItem{}
		err := cursor.Decode(data)
		if err != nil {
			return status.Errorf(
				codes.Internal, fmt.Sprintf("Error while decoding data from MongoDB: %v", err))
		}

		stream.Send(documentToBlog(data))
	}

	if err = cursor.Err(); err != nil {
		return status.Errorf(
			codes.Internal, fmt.Sprintf("Unknown internal error: %v\n", err))
	}
	return nil
}
