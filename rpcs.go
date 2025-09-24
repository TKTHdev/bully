package main

import "fmt"

const (
	ElectionRPC    = "Node.Election"
	CoordinatorRPC = "Node.Coordinator"
	PingRPC        = "Node.Ping"
)

func (n *Node) Election(args *ElectionArgs, reply *ElectionReply) error {
	return nil
}

func (n *Node) Coordinator(args *CoordinatorArgs, reply *CoordinatorReply) error {
	return nil
}

func (n *Node) Ping(args *PingArgs, reply *PingReply) error {
	fmt.Println("ping")
	return nil
}

func (n *Node) sendRPC(targetAddr string, method string, args interface{}, reply interface{}) error {
	client, ok := n.rpcClient[targetAddr]
	if !ok {
		return fmt.Errorf("no RPC client for address: %s", targetAddr)
	}
	if err := client.Call(method, args, reply); err != nil {
		n.rpcClient[targetAddr] = nil
		return err
	}
	return nil
}
