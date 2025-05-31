package main

import (
	pb "github.com/yaninyzwitty/pgxpool-twitter-roach/shared/proto/user"
	"google.golang.org/grpc"
)

type SocialServiceGrpcHandler struct {
	pb.UnimplementedUserServiceServer
}

func NewSocialServiceGrpcHandler(s *grpc.Server) {
	socialServiceGrpcHandler := &SocialServiceGrpcHandler{}
	pb.RegisterUserServiceServer(s, socialServiceGrpcHandler)

}
