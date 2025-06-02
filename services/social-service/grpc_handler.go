package main

import (
	"context"

	pb "github.com/yaninyzwitty/pgxpool-twitter-roach/shared/proto/user"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type SocialServiceGrpcHandler struct {
	pb.UnimplementedUserServiceServer
	repository *PostgresSocialServiceRepository
}

func NewSocialServiceGrpcHandler(s *grpc.Server, repository *PostgresSocialServiceRepository) {
	socialServiceGrpcHandler := &SocialServiceGrpcHandler{
		repository: repository,
	}
	pb.RegisterUserServiceServer(s, socialServiceGrpcHandler)

}

func (s *SocialServiceGrpcHandler) GetUserById(ctx context.Context, req *pb.GetUserByIdRequest) (*pb.GetUserByIdResponse, error) {
	if req.Id <= 0 {
		return nil, status.Errorf(codes.Internal, "Invalid argument: %v", req.Id)
	}

	user, err := s.repository.GetUserById(ctx, uint64(req.Id))
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Error: %v", err)
	}

	return &pb.GetUserByIdResponse{
		User: user,
	}, nil

}
func (s *SocialServiceGrpcHandler) GetUserByEmail(ctx context.Context, req *pb.GetUserByEmailRequest) (*pb.GetUserByEmailResponse, error) {
	if req.Email == "" {
		return nil, status.Errorf(codes.Internal, "Invalid argument: %v", req.Email)
	}

	user, err := s.repository.GetUserByEmail(ctx, req.Email)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Error: %v", err)
	}
	return &pb.GetUserByEmailResponse{
		User: user,
	}, nil

}

func (s *SocialServiceGrpcHandler) GetUsers(ctx context.Context, req *pb.GetUsersRequest) (*pb.GetUsersResponse, error) {
	if req.Limit <= 0 {
		req.Limit = 10
	}

	users, err := s.repository.GetUsers(ctx, int(req.Limit), int(req.Offset))
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Error: %v", err)
	}
	return &pb.GetUsersResponse{
		Users: users,
	}, nil

}

func (s *SocialServiceGrpcHandler) StreamUsers(req *pb.StreamUsersRequest, stream pb.UserService_StreamUsersServer) error {
	ctx := stream.Context()

	limit := int(req.Limit)
	offset := 0
	batchSize := 100
	for {
		users, err := s.repository.GetUsers(ctx, limit, offset)
		if err != nil {
			return status.Errorf(codes.Internal, "failed to fetch users: %v", err)
		}

		if len(users) == 0 {
			break

		}
		for _, user := range users {
			select {
			case <-ctx.Done():
				return ctx.Err()
			default:
				if err := stream.Send(&pb.StreamUsersResponse{
					User: user,
				}); err != nil {
					return status.Errorf(codes.Internal, "failed to send user: %v", err)
				}
			}

		}
		offset += batchSize
		if limit > 0 && offset >= limit {
			break
		}
	}
	return nil
}
