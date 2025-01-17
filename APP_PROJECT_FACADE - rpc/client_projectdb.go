package main

import (
	"fmt"

	"github.com/go-resty/resty/v2"
)

type ProjectDBClient struct {
	url        string
	restClient *resty.Client
}

func NewProjectDBClient(url string) *ProjectDBClient {
	return &ProjectDBClient{
		url:        url,
		restClient: resty.New(),
	}
}

func (p *ProjectDBClient) POST(path string, data map[string]interface{}) (map[string]interface{}, bool) {
	var responseMap map[string]interface{}
	resp, err := p.restClient.R().
		SetHeader("Content-Type", "application/json").
		SetBody(data).
		ForceContentType("application/json").
		SetResult(&responseMap).
		Post(p.url + path)
	fmt.Println("STATUS | POST "+path, resp.Status())

	if err != nil || resp.Status() != "200 OK" {
		fmt.Println("Error:", err)
		return map[string]interface{}{}, false
	}
	return responseMap, true
}
func (p *ProjectDBClient) GET(path string) (map[string]interface{}, bool) {
	var responseMap map[string]interface{}
	resp, err := p.restClient.R().
		ForceContentType("application/json").
		SetResult(&responseMap).
		Get(p.url + path)
	fmt.Println("STATUS | GET"+path, resp.Status())

	if err != nil || resp.Status() != "200 OK" {
		fmt.Println("Error:", err)
		return map[string]interface{}{}, false
	}
	return responseMap, true
}
func (p *ProjectDBClient) GETMULTI(path string) ([]map[string]interface{}, bool) {
	var responseMap []map[string]interface{}
	resp, err := p.restClient.R().
		ForceContentType("application/json").
		SetResult(&responseMap).
		Get(p.url + path)
	fmt.Println("STATUS | GET"+path, resp.Status())

	if err != nil || resp.Status() != "200 OK" {
		fmt.Println("Error:", err)
		return []map[string]interface{}{}, false
	}
	return responseMap, true
}
func (p *ProjectDBClient) DELETE(path string) (map[string]interface{}, bool) {
	var responseMap map[string]interface{}
	resp, err := p.restClient.R().
		ForceContentType("application/json").
		SetResult(&responseMap).
		Delete(p.url + path)
	fmt.Println("STATUS | DELETE"+path, resp.Status())

	if err != nil || resp.Status() != "200 OK" {
		fmt.Println("Error:", err)
		return map[string]interface{}{}, false
	}
	return responseMap, true
}
func (p *ProjectDBClient) DELETE_WITH_BODY(path string, data map[string]interface{}) (map[string]interface{}, bool) {
	var responseMap map[string]interface{}
	resp, err := p.restClient.R().
		SetHeader("Content-Type", "application/json").
		SetBody(data).
		ForceContentType("application/json").
		SetResult(&responseMap).
		Delete(p.url + path)
	fmt.Println("STATUS | POST "+path, resp.Status())

	if err != nil || resp.Status() != "200 OK" {
		fmt.Println("Error:", err)
		return map[string]interface{}{}, false
	}
	return responseMap, true
}
func (p *ProjectDBClient) PUT(path string, data map[string]interface{}) (map[string]interface{}, bool) {
	var responseMap map[string]interface{}
	resp, err := p.restClient.R().
		SetHeader("Content-Type", "application/json").
		SetBody(data).
		ForceContentType("application/json").
		SetResult(&responseMap).
		Put(p.url + path)
	fmt.Println("STATUS | POST "+path, resp.Status())

	if err != nil || resp.Status() != "200 OK" {
		fmt.Println("Error:", err)
		return map[string]interface{}{}, false
	}
	return responseMap, true
}
func (p *ProjectDBClient) PATCH(path string, data map[string]interface{}) (map[string]interface{}, bool) {
	var responseMap map[string]interface{}
	resp, err := p.restClient.R().
		SetHeader("Content-Type", "application/json").
		SetBody(data).
		ForceContentType("application/json").
		SetResult(&responseMap).
		Patch(p.url + path)
	fmt.Println("STATUS | POST "+path, resp.Status())

	if err != nil || resp.Status() != "200 OK" {
		fmt.Println("Error:", err)
		return map[string]interface{}{}, false
	}
	return responseMap, true
}
