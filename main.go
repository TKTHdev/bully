package main

import (
	"fmt"
)

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
