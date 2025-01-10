package main

import (
	"context"
	pb "facade_service/protobuf_generated"
)

type UserServiceServer struct {
	pb.UnimplementedUserServiceServer
	authClient   AuthClient
	userdbClient UserDBClient
}

func (s *UserServiceServer) LoginAndAuthenticate(ctx context.Context, req *pb.LoginRequest) (*pb.AuthResponse, error) {
	// login
	loginResp, ok := s.authClient.Login(req.Email, req.Password)
	if !ok {
		return &pb.AuthResponse{Valid: false}, nil
	}
	token := loginResp["token"].(string)

	// auth
	authResp, err := s.AuthenticateToken(ctx, &pb.TokenRequest{Token: token})
	if err != nil {
		return &pb.AuthResponse{Valid: false}, nil
	}
	return authResp, nil
}

func (s *UserServiceServer) AuthenticateToken(ctx context.Context, req *pb.TokenRequest) (*pb.AuthResponse, error) {
	// validate token
	resp, ok := s.authClient.ValidateToken(req.Token)
	if !ok {
		return &pb.AuthResponse{Valid: false}, nil
	}
	user_id := resp["user_id"].(string)
	// query user info
	resp, ok = s.userdbClient.QueryUser(user_id)
	if !ok {
		return &pb.AuthResponse{Valid: false}, nil
	}
	return &pb.AuthResponse{
		Valid: true,
		Token: req.Token,
		User: &pb.User{
			FirstName: resp["first_name"].(string),
			LastName:  resp["last_name"].(string),
			Id:        resp["id"].(string),
		},
	}, nil
}

func (s *UserServiceServer) CreateAccount(ctx context.Context, req *pb.RegisterRequest) (*pb.AuthResponse, error) {
	// register
	_, ok := s.authClient.Register(req.FirstName, req.LastName, req.Email, req.Password)
	if !ok {
		return &pb.AuthResponse{Valid: false}, nil
	}
	// login & auth
	authResp, err := s.LoginAndAuthenticate(ctx, &pb.LoginRequest{Email: req.Email, Password: req.Password})
	if err != nil {
		return &pb.AuthResponse{Valid: false}, nil
	}
	return authResp, nil
}

func (s *UserServiceServer) GetFriends(ctx context.Context, req *pb.UserID) (*pb.FriendsResponse, error) {
	resp, err := s.GetFriendsOrRequests(req, "friends")
	return resp, err
}

func (s *UserServiceServer) GetFriendRequests(ctx context.Context, req *pb.UserID) (*pb.FriendsResponse, error) {
	resp, err := s.GetFriendsOrRequests(req, "friendRequests")
	return resp, err
}
func (s *UserServiceServer) AddFriend(ctx context.Context, req *pb.FriendEditRequest) (*pb.FriendsResponse, error) {
	resp, err := s.RemoveOrAddFriendsOrRequests(req, "addFriend")
	return resp, err
}
func (s *UserServiceServer) RemoveFriend(ctx context.Context, req *pb.FriendEditRequest) (*pb.FriendsResponse, error) {
	resp, err := s.RemoveOrAddFriendsOrRequests(req, "removeFriend")
	return resp, err
}
func (s *UserServiceServer) AddFriendRequest(ctx context.Context, req *pb.FriendEditRequest) (*pb.FriendsResponse, error) {
	resp, err := s.RemoveOrAddFriendsOrRequests(req, "addFriendRequest")
	return resp, err
}
func (s *UserServiceServer) RemoveFriendRequest(ctx context.Context, req *pb.FriendEditRequest) (*pb.FriendsResponse, error) {
	resp, err := s.RemoveOrAddFriendsOrRequests(req, "removeFriendRequest")
	return resp, err
}

func (s *UserServiceServer) GetFriendsOrRequests(req *pb.UserID, graphqltype string) (*pb.FriendsResponse, error) {
	resp, ok := s.userdbClient.QueryFriendsOrRequests(req.UserId, graphqltype)
	if !ok {
		return &pb.FriendsResponse{Users: []*pb.User{}}, nil
	}
	var users []*pb.User
	for _, userMap := range resp {
		user := &pb.User{
			FirstName: userMap["first_name"].(string),
			LastName:  userMap["last_name"].(string),
			Id:        userMap["id"].(string),
		}
		users = append(users, user)
	}
	return &pb.FriendsResponse{
		Users: users,
	}, nil
}

func (s *UserServiceServer) RemoveOrAddFriendsOrRequests(req *pb.FriendEditRequest, graphqltype string) (*pb.FriendsResponse, error) {
	resp, ok := s.userdbClient.RemoveOrAddFriendsOrRequests(req.UserId, req.FriendId, graphqltype)
	if !ok {
		return &pb.FriendsResponse{Users: []*pb.User{}}, nil
	}
	var users []*pb.User
	for _, userMap := range resp {
		user := &pb.User{
			FirstName: userMap["first_name"].(string),
			LastName:  userMap["last_name"].(string),
			Id:        userMap["id"].(string),
		}
		users = append(users, user)
	}
	return &pb.FriendsResponse{
		Users: users,
	}, nil
}
