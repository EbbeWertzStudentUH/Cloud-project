package main

import (
	"context"
	"log"
	"net"

	pb "facade_service/protobuf_generated"

	"google.golang.org/grpc"
)

type Server struct {
	pb.UnimplementedFacadeServer
}

func (s *Server) Coordinate(ctx context.Context, req *pb.GrpcRequest) (*pb.GrpcResponse, error) {
	log.Printf("Received path segments: %v", req.PathSegments)
	log.Printf("Received query params: %v", req.QueryParams)

	return &pb.GrpcResponse{
		Message: "Processed by Go gRPC Server",
	}, nil
}

func main() {
	listener, err := net.Listen("tcp", ":50052")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterFacadeServer(grpcServer, &Server{})

	log.Println("Go gRPC Server started at :50052")
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
