package main

type (
	ElectionArgs struct {
		addr string
	}
	ElectionReply struct{}

	CoordinatorArgs struct {
		addr string
	}

	CoordinatorReply struct{}

	PingArgs struct {
		addr string
	}

	PingReply struct{}
)
