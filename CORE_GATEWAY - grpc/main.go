package main

import (
	"google.golang.org/grpc"

	pb "facade_service/protobuf_generated"
	"log"
	"net"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file %v", err)
	}

	lis, err := net.Listen("tcp", ":"+os.Getenv("LISTEN_PORT"))
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	auth_client := NewAuthClient(os.Getenv("AUTH_SERVICE_URL"))
	userdb_client := NewUserDBClient(os.Getenv("USERDB_SERVICE_URL"))

	grpcServer := grpc.NewServer()
	pb.RegisterUserServiceServer(grpcServer, &UserServiceServer{authClient: *auth_client, userdbClient: *userdb_client})

	log.Println("Server is running on port " + os.Getenv("LISTEN_PORT"))
	err = grpcServer.Serve(lis)
	if err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}

}
