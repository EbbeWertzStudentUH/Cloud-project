package main

import (
	"log"
	"net/rpc"
	"time"
)

type ProjectFacadeClient struct {
	rpcClient *rpc.Client
}

func NewProjectFacadeClient(url string) *ProjectFacadeClient {
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
	}
}

// args := &common.SomeArgs{A: 5, B: 10}
// 	reply := &common.SomeReply{}

// 	err = client.Call("ProjectService.SomeMethod", args, reply)
