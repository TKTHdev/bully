package main

import "net/rpc"

type Node struct {
	addr      string
	nodeList  []string
	rpcClient map[string]*rpc.Client
}

func NewNode() *Node {
	n := new(Node)
	n.readClusterConfigAndSet("cluster.conf")
	n.readNodeIndexAndSet()
	n.setupRPC()
	return n
}

func (n *Node) run() {
	n.initRPCClients()
	defer n.cleanUpRPCClient()

}
