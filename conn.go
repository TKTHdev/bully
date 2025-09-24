package main

import (
	"fmt"
	"net/rpc"
)

func (n *Node) setupRPC() {
	service := new(Node)
	err := rpc.Register(service)
	if err != nil {
		fmt.Println("Error registering RPC service:", err)
		return
	}
}

func (n *Node) makeRPCClient(targetAddr string) (*rpc.Client, error) {
	client, err := rpc.Dial("tcp", targetAddr)
	if err != nil {
		return nil, err
	}
	return client, nil
}

func (n *Node) initRPCClients() {
	n.rpcClient = make(map[string]*rpc.Client)
	for _, addr := range n.nodeList {
		if addr != n.addr {
			client, err := n.makeRPCClient(addr)
			if err != nil {
				fmt.Println("Error creating RPC client for", addr, ":", err)
			} else {
				n.rpcClient[addr] = client
			}
		}
	}
}

func (n *Node) cleanUpRPCClient() {
	for _, client := range n.rpcClient {
		if client != nil {
			client.Close()
		}
	}
}
