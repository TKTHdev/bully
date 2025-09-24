package main

type (
	ElectionArgs struct {
		Addr string
	}
	ElectionReply struct {
		IsLeader bool
	}

	CoordinatorArgs struct {
		Addr string
	}

	CoordinatorReply struct{}

	PingArgs struct {
		Addr string
	}

	PingReply struct{}
)
