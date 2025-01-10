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

func (n *NotifierClient) SendNotification(user_id string, message string) bool {
	data := map[string]string{
		"messaage": message,
	}
	resp, err := n.restClient.R().
		SetHeader("Content-Type", "application/json").
		SetBody(data).
		ForceContentType("application/json").
		Post(n.url + "/send/notification")
	fmt.Println("STATUS | Notifier service | send notification", resp.Status())

	if err != nil || resp.Status() != "200 OK" {
		fmt.Println("Error:", err)
		return false
	}
	return true
}

func (n *NotifierClient) Subscribe(user_id string, topic_name string, topic_ids []string) bool {
	data := map[string]interface{}{
		"user_id": user_id,
		"topic": map[string]interface{}{
			"name": topic_name,
			"ids":  topic_ids,
		},
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
