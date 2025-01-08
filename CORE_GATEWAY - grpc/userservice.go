package main

import (
	"context"
	pb "facade_service/protobuf_generated"
)

type UserServiceServer struct {
	pb.UnimplementedUserServiceServer
}

func (s *UserServiceServer) LoginAndAuthenticate(ctx context.Context, req *pb.LoginRequest) (*pb.AuthResponse, error) {
	if req.Email == "a@b.com" && req.Password == "pwd" {
		return &pb.AuthResponse{
			Valid:    true,
			Username: "iemand",
		}, nil
	}
	return &pb.AuthResponse{
		Valid:    false,
		Username: "",
	}, nil
}

func (s *UserServiceServer) AuthenticateToken(ctx context.Context, req *pb.TokenRequest) (*pb.AuthResponse, error) {
	if req.Token == "valid-token" {
		return &pb.AuthResponse{
			Valid:    true,
			Username: "iemand",
		}, nil
	}
	return &pb.AuthResponse{
		Valid:    false,
		Username: "",
	}, nil
}

func (s *UserServiceServer) CreateAccount(ctx context.Context, req *pb.RegisterRequest) (*pb.AuthResponse, error) {
	return &pb.AuthResponse{
		Valid:    true,
		Username: req.Username,
	}, nil
}
