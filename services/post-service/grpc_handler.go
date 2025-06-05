package main

import (
	"context"

	pb "github.com/yaninyzwitty/pgxpool-twitter-roach/shared/proto/post"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type PostServiceGrpcHandler struct {
	pb.UnimplementedPostServiceServer

	repository *PostgresPostServiceRepository
}

func NewPostServiceGrpcHandler(s *grpc.Server, repository *PostgresPostServiceRepository) {
	PostServiceGrpcHandler := &PostServiceGrpcHandler{
		repository: repository,
	}
	pb.RegisterPostServiceServer(s, PostServiceGrpcHandler)

}

func (s *PostServiceGrpcHandler) GetPost(ctx context.Context, req *pb.GetPostRequest) (*pb.GetPostResponse, error) {
	if req.PostId <= 0 {
		return nil, status.Errorf(codes.InvalidArgument, "post id is missing")
	}

	post, err := s.repository.GetPost(ctx, req.PostId)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "failed to get comment: %v", err)
	}
	return &pb.GetPostResponse{
		Post: post,
	}, nil
}
