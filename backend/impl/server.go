package main

import (
	"fmt"
	"net"
	"net/rpc"

	magicRPC "github.com/alexanderkogan/magic-bot/backend/rpc"
)

func main() {
	// Start server
	addy, err := net.ResolveTCPAddr("tcp", "0.0.0.0:42586")
	if err != nil {
		panic(fmt.Sprintf("resolve: %s", err))
	}
	magicServer := magicRPC.NewServer()
	if err := rpc.Register(magicServer); err != nil {
		panic(fmt.Sprintf("register: %s", err))
	}

	inbound, err := net.ListenTCP("tcp", addy)
	if err != nil {
		panic(fmt.Sprintf("listen: %s", err))
	}
	rpc.Accept(inbound)

	// V Respond with game state
	// V Start game
	// Tap
	// Move to zone
	// Change life total
}
