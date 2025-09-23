package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
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
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <node_index>")
		return
	}
	index, err := strconv.Atoi(os.Args[1])
	if err != nil || index < 0 || index >= len(n.nodeList) {
		fmt.Println("Invalid node index")
		return
	}
	n.addr = n.nodeList[index]
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
