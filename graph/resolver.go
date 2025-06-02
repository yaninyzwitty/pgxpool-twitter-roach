package graph

import pb "github.com/yaninyzwitty/pgxpool-twitter-roach/shared/proto/user"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	SocialServiceClient pb.UserServiceClient
}
