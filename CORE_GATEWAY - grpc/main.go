package main

import (
	pb "facade_service/protobuf_generated"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	// Listen on a port
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	// Create a new gRPC server
	grpcServer := grpc.NewServer()

	// Register the UserService
	pb.RegisterUserServiceServer(grpcServer, &UserServiceServer{})

	log.Println("Server is running on port 50051...")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
