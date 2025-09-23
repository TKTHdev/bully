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
	addr        string
	processList []string
}

func readClusterConfig(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var processList []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line != "" {
			processList = append(processList, line)
		}
	}
	return processList, scanner.Err()
}

func main() {
	p := &Node{}
	processList, err := readClusterConfig("cluster.conf")
	if err != nil {
		fmt.Println("Error reading cluster config:", err)
		return
	}
	p.processList = processList
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <node_index>")
		return
	}
	index, err := strconv.Atoi(os.Args[1])
	if err != nil || index < 0 || index >= len(processList) {
		fmt.Println("Invalid node index")
		return
	}
	p.addr = processList[index]
	fmt.Println("Node address:", p.addr)
}
