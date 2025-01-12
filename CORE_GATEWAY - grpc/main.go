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

	devstats_client := NewDevstatClient(os.Getenv("SOAP_WSDL_URL"))
	auth_client := NewAuthClient(os.Getenv("AUTH_SERVICE_URL"), devstats_client)
	userdb_client := NewUserDBClient(os.Getenv("USERDB_SERVICE_URL"), devstats_client)
	notifier_client := NewNotifierClient(os.Getenv("NOTIFIER_SERVICE_URL"), devstats_client)

	grpcServer := grpc.NewServer()
	user_service := &UserServiceServer{authClient: *auth_client, userdbClient: *userdb_client, notifierClient: *notifier_client}
	pb.RegisterUserServiceServer(grpcServer, user_service)
	pb.RegisterNotificationServiceServer(grpcServer, &NotificationServiceServer{notifierClient: *notifier_client, userService: user_service})

	log.Println("Server is running on port " + os.Getenv("LISTEN_PORT"))
	err = grpcServer.Serve(lis)
	if err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}

}
