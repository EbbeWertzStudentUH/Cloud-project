package main

import (
	"context"
	"fmt"
	"log"

	"github.com/machinebox/graphql"
)

type User struct {
	Id        string `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}
type GraphQLResponse struct {
	Data struct {
		User User `json:"user"`
	} `json:"data"`
}

type UserDBClient struct {
	graphClient *graphql.Client
}

func NewUserDBClient(url string) *UserDBClient {
	return &UserDBClient{
		graphClient: graphql.NewClient(url),
	}
}

func (u *UserDBClient) Query(query string, vars map[string]interface{}) (*User, bool) {
	req := graphql.NewRequest(query)
	for name, val := range vars {
		req.Var(name, val)
	}
	var response GraphQLResponse
	err := u.graphClient.Run(context.Background(), req, &response)

	if err != nil {
		log.Fatalf("Failed to execute query: %v", err)
		return nil, false
	}
	fmt.Println("STATUS | USERDB service | Query OK")
	return &response.Data.User, true
}
