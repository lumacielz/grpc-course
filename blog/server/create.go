package main

import (
	"context"
	"github.com/lumacielz/grpc-course/blog/proto"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
)

func (s *Server) CreateBlog(ctx context.Context, in *proto.Blog) (*proto.BlogId, error) {
	log.Printf("Create was invoked with %v\n", in)
	data := BlogItem{
		ID:       primitive.NewObjectID(),
		AuthorId: in.AuthorId,
		Title:    in.Title,
		Content:  in.Content,
	}

	res, err := collection.InsertOne(ctx, data)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	oid, ok := res.InsertedID.(primitive.ObjectID)
	if !ok {
		return nil, status.Errorf(codes.Internal, "Cannot convert to ObjectID")
	}

	return &proto.BlogId{
		Id: oid.Hex(),
	}, nil
}
