package main

import (
	"fmt"

	"github.com/go-resty/resty/v2"
)

type UserDBClient struct {
	url        string
	restClient *resty.Client
}

func NewUserDBClient(url string) *UserDBClient {
	return &UserDBClient{
		url:        url,
		restClient: resty.New(),
	}
}

func (u *UserDBClient) QueryUSernames(user_id string) (map[string]interface{}, bool) {
	query := `query($id: String!) {
  		user(id: $id) {
    		first_name
    		last_name
  		}
	}`
	vars := map[string]interface{}{"id": user_id}
	resp, ok := u.Query(query, vars)
	return resp["user"], ok
}

func (a *UserDBClient) Query(query string, vars map[string]interface{}) (map[string]map[string]interface{}, bool) {
	data := map[string]interface{}{
		"query":     query,
		"variables": vars,
	}
	var responseMap map[string]map[string]map[string]interface{}
	resp, err := a.restClient.R().
		SetHeader("Content-Type", "application/json").
		SetBody(data).
		ForceContentType("application/json").
		SetResult(&responseMap).
		Post(a.url + "/users")
	fmt.Println("STATUS | UserDB service | Query User", resp.Status())

	if err != nil || resp.Status() != "200 OK" {
		fmt.Println("Error:", err)
		return map[string]map[string]interface{}{}, false
	}
	return responseMap["data"], true
}
