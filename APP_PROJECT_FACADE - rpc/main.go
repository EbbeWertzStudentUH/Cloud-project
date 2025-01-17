package main

import (
	"log"
	"net/http"
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
	rpc.Register(service)
	rpc.HandleHTTP()
	log.Println("listening on " + "0.0.0.0:" + os.Getenv("LISTEN_PORT"))
	err = http.ListenAndServe(":"+os.Getenv("LISTEN_PORT"), nil)
	if err != nil {
		log.Fatal("Error starting the RPC server: ", err)
	}
}
