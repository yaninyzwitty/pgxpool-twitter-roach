package main

import (
	"context"

	pb "github.com/yaninyzwitty/pgxpool-twitter-roach/shared/proto/comment"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type CommentServiceGrpcHandler struct {
	pb.UnimplementedCommentServiceServer

	repository *PostgresCommentServiceRepository
}

func NewCommentServiceGrpcHandler(s *grpc.Server, repository *PostgresCommentServiceRepository) {
	commentServiceGrpcHandler := &CommentServiceGrpcHandler{
		repository: repository,
	}
	pb.RegisterCommentServiceServer(s, commentServiceGrpcHandler)

}

func (s *CommentServiceGrpcHandler) GetComment(ctx context.Context, req *pb.GetCommentRequest) (*pb.GetCommentResponse, error) {
	if req.CommentId <= 0 {
		return nil, status.Errorf(codes.InvalidArgument, "comment id is required")
	}

	comment, err := s.repository.GetComment(ctx, req.CommentId)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to find comment: %v", err)
	}

	return &pb.GetCommentResponse{
		Comment: comment,
	}, nil
}
