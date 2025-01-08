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
	resp, ok = s.userdbClient.QueryUSernames(user_id)
	if !ok {
		return &pb.AuthResponse{Valid: false}, nil
	}
	return &pb.AuthResponse{
		Valid:     true,
		Token:     req.Token,
		FirstName: resp["first_name"].(string),
		LastName:  resp["last_name"].(string),
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
