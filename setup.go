package main

import (
	"bufio"
	"fmt"
	"net/rpc"
	"os"
	"strconv"
	"strings"
)

func (n *Node) readClusterConfigAndSet(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line != "" {
			n.nodeList = append(n.nodeList, line)
		}
	}
}

func (n *Node) readNodeIndexAndSet() {
	if len(os.Args) < 1 {
		fmt.Println("Usage: go run main.go <node_index>")
		return
	}
	index, err := strconv.Atoi(os.Args[1])
	if err != nil || index < -1 || index >= len(n.nodeList) {
		fmt.Println(err)
		return
	}
	n.addr = n.nodeList[index]
}

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
