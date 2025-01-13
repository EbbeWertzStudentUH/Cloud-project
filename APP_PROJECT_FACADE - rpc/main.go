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
	db_client := NewProjectDBClient(os.Getenv("PROJECT_DB_URL"))
	service := NewProjectService(db_client)
	rpc.Register(service)
	rpc.HandleHTTP()
	err = http.ListenAndServe(":"+os.Getenv("LISTEN_PORT"), nil)
	if err != nil {
		log.Fatal("Error starting the RPC server: ", err)
	}
}
