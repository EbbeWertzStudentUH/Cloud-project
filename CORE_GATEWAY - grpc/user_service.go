package main

import (
	"context"
	pb "facade_service/protobuf_generated"
)

// UserServiceServer implementation
type UserServiceServer struct {
	pb.UnimplementedUserServiceServer
}

// LoginAndAuthenticate implementation
func (s *UserServiceServer) LoginAndAuthenticate(ctx context.Context, req *pb.LoginRequest) (*pb.AuthResponse, error) {
	// Mock implementation
	if req.Email == "test@example.com" && req.Password == "password" {
		return &pb.AuthResponse{
			Valid:    true,
			Username: "test_user",
		}, nil
	}
	return &pb.AuthResponse{
		Valid:    false,
		Username: "",
	}, nil
}

// AuthenticateToken implementation
func (s *UserServiceServer) AuthenticateToken(ctx context.Context, req *pb.TokenRequest) (*pb.AuthResponse, error) {
	// Mock token validation
	if req.Token == "valid-token" {
		return &pb.AuthResponse{
			Valid:    true,
			Username: "test_user",
		}, nil
	}
	return &pb.AuthResponse{
		Valid:    false,
		Username: "",
	}, nil
}

// CreateAccount implementation
func (s *UserServiceServer) CreateAccount(ctx context.Context, req *pb.RegisterRequest) (*pb.AuthResponse, error) {
	// Mock account creation
	return &pb.AuthResponse{
		Valid:    true,
		Username: req.Username,
	}, nil
}
