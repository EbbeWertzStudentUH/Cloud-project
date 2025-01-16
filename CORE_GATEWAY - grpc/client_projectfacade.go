package main

import (
	"log"
	"net/rpc"
	"time"
)

type ProjectFacadeClient struct {
	rpcClient *rpc.Client
	dsc       *DevstatClient
}

func NewProjectFacadeClient(url string, dsc *DevstatClient) *ProjectFacadeClient {
	var client *rpc.Client
	var err error
	for {
		client, err = rpc.Dial("tcp", url)
		if err == nil {
			log.Println("connected to rpc:", url)
			break
		}
		log.Println("could not connect to gRPC server. Trying again in 3s")
		time.Sleep(3 * time.Second)
	}
	return &ProjectFacadeClient{
		rpcClient: client,
		dsc:       dsc,
	}
}

func (p *ProjectFacadeClient) call(function string, req any, resp any) any {
	dsc_id := p.dsc.Start("Project facade", "(Go)RPC", function)
	p.rpcClient.Call("ProjectService."+function, req, resp)
	p.dsc.End(dsc_id)
	return resp
}
