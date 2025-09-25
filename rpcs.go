package main

import "fmt"

const (
	ElectionRPC    = "Node.Election"
	CoordinatorRPC = "Node.Coordinator"
	PingRPC        = "Node.Ping"
)

func (n *Node) Election(args *ElectionArgs, reply *ElectionReply) error {
	fmt.Println("I am a node: ", n.addr, " received Election from ", args.Addr, " isLeader:", n.isLeader)
	reply.IsLeader = n.isLeader
	if !n.isLeader {
		n.leader = ""
		n.startElection()
	}
	return nil
}

func (n *Node) Coordinator(args *CoordinatorArgs, reply *CoordinatorReply) error {
	fmt.Println("Received Coordinator from ", args.Addr)
	n.leader = args.Addr
	n.isLeader = false
	return nil
}

func (n *Node) Ping(args *PingArgs, reply *PingReply) error {
	fmt.Println("ping")
	return nil
}

func (n *Node) sendRPC(targetAddr string, method string, args interface{}, reply interface{}) error {
	n.mu.Lock()
	defer n.mu.Unlock()
	client, ok := n.rpcClient[targetAddr]
	if !ok {
		return fmt.Errorf("no RPC client for address: %s", targetAddr)
	}
	if err := client.Call(method, args, reply); err != nil {
		n.rpcClient[targetAddr].Close()
		delete(n.rpcClient, targetAddr)
		return err
	}
	return nil
}
