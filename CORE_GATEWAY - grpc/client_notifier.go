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
	data := map[string]interface{}{
		"user_id": user_id,
		"notification": map[string]interface{}{
			"message": message,
		},
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

func (n *NotifierClient) SendUpdate(user_id string, update_type string, subject string, update_data map[string]interface{}) bool {
	data := map[string]interface{}{
		"user_id": user_id,
		"update": map[string]interface{}{
			"type":    update_type,
			"subject": subject,
			"data":    update_data,
		},
	}
	resp, err := n.restClient.R().
		SetHeader("Content-Type", "application/json").
		SetBody(data).
		ForceContentType("application/json").
		Post(n.url + "/send/update")
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
func (n *NotifierClient) UnSubscribe(user_id string, topic_name string, topic_ids []string) bool {
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
		Delete(n.url + "/subscribe")
	fmt.Println("STATUS | Notifier service | subscribe", resp.Status())

	if err != nil || resp.Status() != "200 OK" {
		fmt.Println("Error:", err)
		return false
	}
	return true
}
