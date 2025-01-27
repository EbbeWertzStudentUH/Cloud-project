package main

import (
	"fmt"
	"time"

	"github.com/go-resty/resty/v2"
)

type UserDBClient struct {
	url        string
	restClient *resty.Client
}

func NewUserDBClient(url string) *UserDBClient {
	client := resty.New()
	for {
		_, err := client.R().Head(url)
		if err == nil {
			fmt.Println("connected to GRAPQL User DB.")
			break
		}
		fmt.Println("could not connect to GRAPHQL server: user db. Trying again in 3s")
		time.Sleep(3 * time.Second)
	}
	return &UserDBClient{
		url:        url,
		restClient: resty.New(),
	}
}

func (a *UserDBClient) QueryUsers(ids []string) ([]map[string]interface{}, bool) {
	query := `query($ids: [String!]!) {
  		users(ids: $ids) {
    		id
    		first_name
    		last_name
  		}
	}`
	vars := map[string][]string{"ids": ids}
	data := map[string]interface{}{
		"query":     query,
		"variables": vars,
	}
	var responseMap map[string]map[string][]map[string]interface{}
	resp, err := a.restClient.R().
		SetHeader("Content-Type", "application/json").
		SetBody(data).
		ForceContentType("application/json").
		SetResult(&responseMap).
		Post(a.url + "/users")
	fmt.Println("STATUS | UserDB service | Query Users", resp.Status())
	if err != nil || resp.Status() != "200 OK" {
		fmt.Println("Error:", err)
		return []map[string]interface{}{}, false
	}
	return responseMap["data"]["users"], true
}
