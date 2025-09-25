package main

import (
	"fmt"
	"net/rpc"
	"slices"
	"sync"
	"time"
)

type Node struct {
	addr      string
	nodeList  []string
	rpcClient map[string]*rpc.Client
	isLeader  bool
	leader    string
	mu        sync.Mutex
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

func (n *Node) electionToUpperNodes() bool {

	for _, addr := range n.nodeList {
		if addr > n.addr {
			args := &ElectionArgs{Addr: n.addr}
			reply := &ElectionReply{}
			err := n.sendRPC(addr, ElectionRPC, args, reply)
			fmt.Println("reply from", addr, ":", reply)
			if err != nil {
				fmt.Println("Error sending Election RPC to", addr, ":", err)
				continue
			} else {
				if reply.IsLeader {
					n.leader = addr
				}
				return true
			}
		}
	}
	return false
}

func (n *Node) coordinatorToLowerNodes() {
	for _, addr := range n.nodeList {
		if addr < n.addr {
			args := &CoordinatorArgs{Addr: n.addr}
			reply := &CoordinatorReply{}
			err := n.sendRPC(addr, CoordinatorRPC, args, reply)
			if err != nil {
				fmt.Println("Error sending Coordinator RPC to", addr, ":", err)
				continue
			} else {
				fmt.Println("Sent Coordinator RPC to", addr)
			}
		}
	}
}

func (n *Node) pingLeader() bool {
	if n.leader == "" || n.isLeader {
		return false
	}
	args := &PingArgs{Addr: n.addr}
	reply := &PingReply{}
	err := n.sendRPC(n.leader, PingRPC, args, reply)
	if err != nil {
		return false
	} else {
		return true
	}
}

func (n *Node) startElection() {
	if slices.Max(n.nodeList) == n.addr {
		n.isLeader = true
		n.leader = n.addr
		n.coordinatorToLowerNodes()
		return
	}

	respFromUpperNode := n.electionToUpperNodes()
	if !respFromUpperNode {
		fmt.Println("I am a leader")
		n.isLeader = true
		n.leader = n.addr
		n.coordinatorToLowerNodes()
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

		fmt.Println("leader:", n.leader)

		ping := n.pingLeader()
		fmt.Println("Ping leader:", ping)
		if ping {
			continue
		} else {
			n.startElection()
		}
	}
}
