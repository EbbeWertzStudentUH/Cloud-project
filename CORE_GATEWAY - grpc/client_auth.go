package main

import (
	"fmt"

	"github.com/go-resty/resty/v2"
)

type AuthClient struct {
	url        string
	restClient *resty.Client
	dsc        *DevstatClient
}

func NewAuthClient(url string, dsc *DevstatClient) *AuthClient {
	return &AuthClient{
		url:        url,
		restClient: resty.New(),
		dsc:        dsc,
	}
}

func (a *AuthClient) Register(first_name string, last_name string, email string, pwd string) (map[string]interface{}, bool) {
	dsc_id := a.dsc.Start("Auth", "REST", "Register")
	data := map[string]string{
		"first_name": first_name,
		"last_name":  last_name,
		"email":      email,
		"password":   pwd,
	}
	var responseMap map[string]interface{}
	resp, err := a.restClient.R().
		SetHeader("Content-Type", "application/json").
		SetBody(data).
		ForceContentType("application/json").
		SetResult(&responseMap).
		Post(a.url + "/register")
	fmt.Println("STATUS | Auth service | Register", resp.Status())

	if err != nil || resp.Status() != "200 OK" {
		fmt.Println("Error:", err)
		return map[string]interface{}{}, false
	}
	a.dsc.End(dsc_id)
	return responseMap, true
}

func (a *AuthClient) Login(email string, pwd string) (map[string]interface{}, bool) {
	dsc_id := a.dsc.Start("Auth", "REST", "Login")
	data := map[string]string{
		"email":    email,
		"password": pwd,
	}
	var responseMap map[string]interface{}
	resp, err := a.restClient.R().
		SetHeader("Content-Type", "application/json").
		SetBody(data).
		ForceContentType("application/json").
		SetResult(&responseMap).
		Post(a.url + "/login")
	fmt.Println("STATUS | Auth service | Login", resp.Status())

	if err != nil || resp.Status() != "200 OK" {
		fmt.Println("Error:", err)
		return map[string]interface{}{}, false
	}
	a.dsc.End(dsc_id)
	return responseMap, true
}

func (a *AuthClient) ValidateToken(token string) (map[string]interface{}, bool) {
	dsc_id := a.dsc.Start("Auth", "REST", "ValidateToken")
	var responseMap map[string]interface{}
	resp, err := a.restClient.R().
		SetHeader("Authorization", "Bearer "+token).
		ForceContentType("application/json").
		SetResult(&responseMap).
		Get(a.url + "/verify_token")
	fmt.Println("STATUS | Auth service | Verify Token", resp.Status())

	if err != nil || resp.Status() != "200 OK" || responseMap["valid"] == false {
		fmt.Println("Error:", err)
		return map[string]interface{}{}, false
	}
	a.dsc.End(dsc_id)
	return responseMap, true
}
