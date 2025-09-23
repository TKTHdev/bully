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
	addr        string
	processList []string
}

func main() {
	fmt.Println("bully")
}
