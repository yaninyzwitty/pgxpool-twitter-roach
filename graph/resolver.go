package graph

import (
	commentpb "github.com/yaninyzwitty/pgxpool-twitter-roach/shared/proto/comment"
	pb "github.com/yaninyzwitty/pgxpool-twitter-roach/shared/proto/user"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	SocialServiceClient  pb.UserServiceClient
	CommentServiceClient commentpb.CommentServiceClient
}
