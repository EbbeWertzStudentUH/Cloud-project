package main

import (
	"fmt"

	"github.com/go-resty/resty/v2"
)

type NotifierClient struct {
	url        string
	restClient *resty.Client
}

func NewNotifierClient(url string) *NotifierClient {
	return &NotifierClient{
		url:        url,
		restClient: resty.New(),
	}
}

func (n *NotifierClient) Subscribe(user_id string, topic_name string, topic_ids []string) bool {
	data := map[string]interface{}{
		"user_id":    user_id,
		"topic_name": topic_name,
		"topic_ids":  topic_ids,
	}
	resp, err := n.restClient.R().
		SetHeader("Content-Type", "application/json").
		SetBody(data).
		ForceContentType("application/json").
		Post(n.url + "/subscribe")
	fmt.Println("STATUS | Notifier service | subscribe", resp.Status())

	if err != nil || resp.Status() != "200 OK" {
		fmt.Println("Error:", err)
		return false
	}
	return true
}
