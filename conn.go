package main

import (
	"fmt"
	"net"
	"net/rpc"
)

func (n *Node) setupRPCListen() {
	err := rpc.Register(n)
	if err != nil {
		fmt.Println("Error registering RPC service:", err)
		return
	}
	l, err := net.Listen("tcp", n.addr)
	if err != nil {
		fmt.Println("Error starting RPC listener:", err)
		return
	}
	fmt.Println("Listening for RPCs on", n.addr)
	defer l.Close()
	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Error accepting RPC connection:", err)
			continue
		}
		go rpc.ServeConn(conn)
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
	for {
		for _, addr := range n.nodeList {
			if addr != n.addr {
				if _, exists := n.rpcClient[addr]; !exists {
					client, err := n.makeRPCClient(addr)
					if err != nil {
						//fmt.Println("Error connecting to", addr, ":", err)
						continue
					}
					n.rpcClient[addr] = client
					fmt.Println("Connected to", addr)
				}
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
