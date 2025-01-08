package main

import (
	"context"
	pb "facade_service/protobuf_generated"
)

type UserServiceServer struct {
	pb.UnimplementedUserServiceServer
	authClient AuthClient
}

func (s *UserServiceServer) LoginAndAuthenticate(ctx context.Context, req *pb.LoginRequest) (*pb.AuthResponse, error) {
	resp, err := s.authClient.Login(req.Email, req.Password)
	if err != nil {
		return &pb.AuthResponse{Valid: false}, nil
	}
	token := resp["token"].(string)
	print("token: " + token)
	return &pb.AuthResponse{
		Valid:     true,
		FirstName: "dummy_name",
		LastName:  "dummy_name",
	}, nil
}

func (s *UserServiceServer) AuthenticateToken(ctx context.Context, req *pb.TokenRequest) (*pb.AuthResponse, error) {
	resp, err := s.authClient.ValidateToken(req.Token)
	if err != nil {
		return &pb.AuthResponse{Valid: false}, nil
	}
	user_id := resp["user_id"].(string)
	print("userid: " + user_id)
	return &pb.AuthResponse{
		Valid:     true,
		FirstName: "dummy_name",
		LastName:  "dummy_name",
	}, nil
}

func (s *UserServiceServer) CreateAccount(ctx context.Context, req *pb.RegisterRequest) (*pb.AuthResponse, error) {
	resp, ok := s.authClient.Register(req.FirstName, req.LastName, req.Email, req.Password)
	if !ok {
		return &pb.AuthResponse{Valid: false}, nil
	}
	user := resp["user"].(map[string]interface{})
	return &pb.AuthResponse{
		Valid:     true,
		FirstName: user["first_name"].(string),
		LastName:  user["last_name"].(string),
	}, nil
}
