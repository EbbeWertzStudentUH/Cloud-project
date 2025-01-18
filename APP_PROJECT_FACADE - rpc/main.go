package main

import (
	"log"
	"net"
	"net/rpc"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file %v", err)
	}
	userdb_client := NewUserDBClient(os.Getenv("USERDB_SERVICE_URL"))
	projectdb_client := NewProjectDBClient(os.Getenv("PROJECT_DB_URL"))
	service := NewProjectService(projectdb_client, userdb_client)
	err = rpc.Register(service)
	if err != nil {
		log.Fatal("Error registering RPC service:", err)
	}
	service.GetFullProjectById(&GetProjectByIdRequest{Proj_id: "124f84b3-b0d1-4acf-83b9-1445c1d08ab7"}, &FullProject{})

	listener, err := net.Listen("tcp", ":"+os.Getenv("LISTEN_PORT"))
	if err != nil {
		log.Fatalf("Error starting TCP server: %v", err)
	}
	log.Println("listening on " + "0.0.0.0:" + os.Getenv("LISTEN_PORT"))

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("Connection error: %v", err)
			continue
		}
		go rpc.ServeConn(conn)
	}
}
