package main

import (
	"fmt"

	"github.com/go-resty/resty/v2"
)

type NotifierClient struct {
	url        string
	restClient *resty.Client
	dsc        *DevstatClient
}

func NewNotifierClient(url string, dsc *DevstatClient) *NotifierClient {
	return &NotifierClient{
		url:        url,
		restClient: resty.New(),
		dsc:        dsc,
	}
}

func (n *NotifierClient) SendNotification(user_id string, message string) bool {
	dsc_id := n.dsc.Start("Notifier", "REST", "Send Notification")
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
	n.dsc.End(dsc_id)
	return true
}

func (n *NotifierClient) SendUpdate(user_id string, update_type string, subject string, update_data map[string]interface{}) bool {
	dsc_id := n.dsc.Start("Notifier", "REST", "Send Update")

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
	n.dsc.End(dsc_id)
	return true
}

func (n *NotifierClient) PublishNotification(topic_name string, topic_id string, message string) bool {
	dsc_id := n.dsc.Start("Notifier", "REST", "Publish Notification")
	data := map[string]interface{}{
		"topic": map[string]interface{}{
			"name": topic_name,
			"ids":  []string{topic_id},
		},
		"notification": map[string]interface{}{
			"message": message,
		},
	}
	resp, err := n.restClient.R().
		SetHeader("Content-Type", "application/json").
		SetBody(data).
		ForceContentType("application/json").
		Post(n.url + "/publish/notification")
	fmt.Println("STATUS | Notifier service | send notification", resp.Status())

	if err != nil || resp.Status() != "200 OK" {
		fmt.Println("Error:", err)
		return false
	}
	n.dsc.End(dsc_id)
	return true
}

func (n *NotifierClient) PublishUpdate(topic_name string, topic_id string, update_type string, subject string, update_data map[string]interface{}) bool {
	dsc_id := n.dsc.Start("Notifier", "REST", "Publish Update")

	data := map[string]interface{}{
		"topic": map[string]interface{}{
			"name": topic_name,
			"ids":  []string{topic_id},
		},
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
		Post(n.url + "/publish/update")
	fmt.Println("STATUS | Notifier service | send notification", resp.Status())

	if err != nil || resp.Status() != "200 OK" {
		fmt.Println("Error:", err)
		return false
	}
	n.dsc.End(dsc_id)
	return true
}

func (n *NotifierClient) Subscribe(user_id string, topic_name string, topic_ids []string) bool {
	dsc_id := n.dsc.Start("Notifier", "REST", "Subscribe")

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
	n.dsc.End(dsc_id)
	return true
}
func (n *NotifierClient) UnSubscribe(user_id string, topic_name string, topic_ids []string) bool {
	dsc_id := n.dsc.Start("Notifier", "REST", "UnSubscribe")

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
	n.dsc.End(dsc_id)
	return true
}
