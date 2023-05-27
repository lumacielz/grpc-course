package main

import (
	"context"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/lumacielz/grpc-course/blog/proto"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"log"
)

func (s *Server) UpdateBlog(ctx context.Context, in *proto.Blog) (*empty.Empty, error) {
	log.Printf("UpdateBlog was invoked with %v\n", in)

	oid, err := primitive.ObjectIDFromHex(in.Id)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "Cannot parse ID")
	}

	data := BlogItem{
		AuthorId: in.AuthorId,
		Title:    in.Title,
		Content:  in.Content,
	}

	res, err := collection.UpdateOne(ctx, bson.M{"_id": oid}, bson.M{"$set": &data})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Cannot update")
	}

	if res.MatchedCount == 0 {
		return nil, status.Errorf(codes.NotFound, "Cannot find blog with id")
	}
	return &emptypb.Empty{}, nil
}
