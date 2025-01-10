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

func (u *UserDBClient) QueryUser(user_id string) (map[string]interface{}, bool) {
	query := `query($id: String!) {
  		user(id: $id) {
			id
    		first_name
    		last_name
  		}
	}`
	vars := map[string]interface{}{"id": user_id}
	resp, ok := u.QuerySingle(query, vars)
	return resp["user"], ok
}
func (u *UserDBClient) QueryFriendsOrRequests(user_id string, graphql_type string) ([]map[string]interface{}, bool) {
	query := `query($id: String!) {
  		` + graphql_type + `(id: $id) {
			id
    		first_name
    		last_name
  		}
	}`
	vars := map[string]interface{}{"id": user_id}
	resp, ok := u.QueryMultiple(query, vars)
	return resp[graphql_type], ok
}
func (u *UserDBClient) RemoveOrAddFriendsOrRequests(user_id string, friend_id, graphql_type string) ([]map[string]interface{}, bool) {
	query := `mutation($user_id: String!, $friend_id: String!) {
  		` + graphql_type + `(user_id: $user_id, friend_id: $friend_id) {
			id
    		first_name
    		last_name
  		}
	}`
	vars := map[string]interface{}{"user_id": user_id, "friend_id": friend_id}
	resp, ok := u.QueryMultiple(query, vars)
	return resp[graphql_type], ok
}

func (a *UserDBClient) QuerySingle(query string, vars map[string]interface{}) (map[string]map[string]interface{}, bool) {
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
func (a *UserDBClient) QueryMultiple(query string, vars map[string]interface{}) (map[string][]map[string]interface{}, bool) {
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
	fmt.Println("STATUS | UserDB service | Query User", resp.Status())

	if err != nil || resp.Status() != "200 OK" {
		fmt.Println("Error:", err)
		return map[string][]map[string]interface{}{}, false
	}
	return responseMap["data"], true
}
