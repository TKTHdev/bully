package main

import (
	"fmt"
)

type (
	ElectionArgs  struct{}
	ElectionReply struct{}

	CoordinatorArgs struct{}

	CoordinatorReply struct{}

	PingArgs struct{}

	PingReply struct{}
)

// Node struct
type Node struct {
	addr     string
	nodeList []string
}

func NewNode() *Node {
	n := new(Node)
	n.readClusterConfigAndSet("cluster.conf")
	n.readNodeIndexAndSet()
	return n
}

func main() {
	n := NewNode()
	fmt.Println("Node address:", n.addr)
	fmt.Println("Cluster nodes:", n.nodeList)
}
