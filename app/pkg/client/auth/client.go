package auth_client

import (
	"context"
	"math/rand"

	"github.com/ducnd58233/go-newsfeed-grpc/app/pkg/types/proto/pb/auth_pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type randomClient struct {
	clients []auth_pb.AuthClient
}

// EditUser implements auth_pb.AuthClient.
func (a *randomClient) EditUser(ctx context.Context, in *auth_pb.EditUserRequest, opts ...grpc.CallOption) (*auth_pb.EditUserResponse, error) {
	return a.clients[rand.Intn(len(a.clients))].EditUser(ctx, in, opts...)
}

func (a *randomClient) CheckUserAuthentication(ctx context.Context, in *auth_pb.CheckUserAuthenticationRequest, opts ...grpc.CallOption) (*auth_pb.CheckUserAuthenticationResponse, error) {
	return a.clients[rand.Intn(len(a.clients))].CheckUserAuthentication(ctx, in, opts...)
}

func (a *randomClient) CreateUser(ctx context.Context, in *auth_pb.UserDetailInfo, opts ...grpc.CallOption) (*auth_pb.UserResult, error) {
	return a.clients[rand.Intn(len(a.clients))].CreateUser(ctx, in, opts...)
}

func (a *randomClient) GetUserFollower(ctx context.Context, in *auth_pb.UserInfo, opts ...grpc.CallOption) (*auth_pb.UserFollower, error) {
	return a.clients[rand.Intn(len(a.clients))].GetUserFollower(ctx, in, opts...)
}

func (a *randomClient) FollowUser(ctx context.Context, in *auth_pb.FollowUserRequest, opts ...grpc.CallOption) (*auth_pb.FollowUserResponse, error) {
	return a.clients[rand.Intn(len(a.clients))].FollowUser(ctx, in, opts...)
}

func (a *randomClient) UnfollowUser(ctx context.Context, in *auth_pb.UnfollowUserRequest, opts ...grpc.CallOption) (*auth_pb.UnfollowUserResponse, error) {
	return a.clients[rand.Intn(len(a.clients))].UnfollowUser(ctx, in, opts...)
}

func (a *randomClient) GetFollowerList(ctx context.Context, in *auth_pb.GetFollowerListRequest, opts ...grpc.CallOption) (*auth_pb.GetFollowerListResponse, error) {
	return a.clients[rand.Intn(len(a.clients))].GetFollowerList(ctx, in, opts...)
}
func NewClient(hosts []string) (auth_pb.AuthClient, error) {
	var opts = []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	clients := make([]auth_pb.AuthClient, 0, len(hosts))
	for _, host := range hosts {
		conn, err := grpc.Dial(host, opts...)
		if err != nil {
			return nil, err
		}
		client := auth_pb.NewAuthClient(conn)
		clients = append(clients, client)
	}
	return &randomClient{clients}, nil
}
