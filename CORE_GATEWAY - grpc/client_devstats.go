package main

import (
	"facade_service/soap_generated"

	"github.com/hooklift/gowsdl/soap"
)

type DevstatClient struct {
	soapClient soap_generated.Main
}

func NewDevstatClient(url string) *DevstatClient {
	client := soap.NewClient(url)
	return &DevstatClient{
		soapClient: soap_generated.NewMain(client),
	}
}

func (d *DevstatClient) Start(serviceName string, serviceType string, identifier string) string {
	resp, _ := d.soapClient.RegisterOutgoingStart(&soap_generated.RegisterOutgoingStart{
		ServiceType: serviceType,
		Identifier:  identifier,
		ServiceName: serviceName,
	})
	return resp.Return_
}

func (d *DevstatClient) End(requestId string) {
	d.soapClient.RegisterOutgoingEnd(&soap_generated.RegisterOutgoingEnd{
		RequestId: requestId,
	})
}
