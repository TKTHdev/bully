package main

import (
	"fmt"
	"net/rpc"
	"time"
)

type Node struct {
	addr      string
	nodeList  []string
	rpcClient map[string]*rpc.Client
	isLeader  bool
	leader    string
}

func NewNode() *Node {
	n := new(Node)
	n.readClusterConfigAndSet("cluster.conf")
	n.readNodeIndexAndSet()
	n.isLeader = false
	n.leader = ""
	go n.setupRPCListen()
	return n
}

func (n *Node) PingToUpperNodes() {
	for _, addr := range n.nodeList {
		if addr > n.addr {
			args := &PingArgs{}
			reply := &PingReply{}
			err := n.sendRPC(addr, PingRPC, args, reply)
			if err != nil {

			} else {
				continue
			}
		}
	}

}

func (n *Node) run() {
	go n.initRPCClients()
	defer n.cleanUpRPCClient()
	fmt.Println("Process: ", n.addr, "started")

	for {
		time.Sleep(100 * time.Millisecond)
		if n.isLeader {
			continue
		}

	}
}
