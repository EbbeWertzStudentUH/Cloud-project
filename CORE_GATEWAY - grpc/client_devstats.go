package main

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/go-resty/resty/v2"
)

const soapStartRequest = `
<soapenv:Envelope xmlns:soapenv="http://schemas.xmlsoap.org/soap/envelope/" xmlns:web="http://ebbew.be/">
    <soapenv:Header/>
    <soapenv:Body>
        <web:registerOutgoingStart>
            <serviceType>%s</serviceType>
            <serviceName>%s</serviceName>
            <identifier>%s</identifier>
        </web:registerOutgoingStart>
    </soapenv:Body>
</soapenv:Envelope>`

const soapEndRequest = `
<soapenv:Envelope xmlns:soapenv="http://schemas.xmlsoap.org/soap/envelope/" xmlns:web="http://ebbew.be/">
    <soapenv:Header/>
    <soapenv:Body>
        <web:registerOutgoingEnd>
            <requestId>%s</requestId>
        </web:registerOutgoingEnd>
    </soapenv:Body>
</soapenv:Envelope>`

type DevstatClient struct {
	url        string
	restClient *resty.Client
}

func NewDevstatClient(url string) *DevstatClient {
	return &DevstatClient{
		url:        url,
		restClient: resty.New(),
	}
}

func (d *DevstatClient) Start(serviceName string, serviceType string, identifier string) string {
	request := fmt.Sprintf(soapStartRequest, serviceType, serviceName, identifier)
	resp := sendSOAPRequest(d.url, request)
	requestId := extractReturn(resp)
	return requestId
}

func (d *DevstatClient) End(requestId string) {
	request := fmt.Sprintf(soapEndRequest, requestId)
	sendSOAPRequest(d.url, request)
}

func sendSOAPRequest(url string, request string) string {
	req, _ := http.NewRequest("POST", url, bytes.NewBuffer([]byte(request)))
	req.Header.Set("Content-Type", "text/xml; charset=utf-8")
	req.Header.Set("SOAPAction", `"#POST"`)
	client := &http.Client{}
	resp, _ := client.Do(req)
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	return string(body)
}

type Envelope struct {
	XMLName xml.Name `xml:"Envelope"`
	Body    Body     `xml:"Body"`
}

type Body struct {
	XMLName                       xml.Name                      `xml:"Body"`
	RegisterOutgoingStartResponse RegisterOutgoingStartResponse `xml:"registerOutgoingStartResponse"`
}

type RegisterOutgoingStartResponse struct {
	XMLName xml.Name `xml:"registerOutgoingStartResponse"`
	Return  string   `xml:"return"`
}

func extractReturn(xmlResponse string) string {
	var envelope Envelope
	xml.Unmarshal([]byte(xmlResponse), &envelope)
	return envelope.Body.RegisterOutgoingStartResponse.Return
}
