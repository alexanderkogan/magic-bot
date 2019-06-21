package rpc

import (
	"fmt"
	"net/rpc"

	"github.com/alexanderkogan/magic-bot/backend"
)

// This is a wrapper for the RPC client, so that the magic client can use the more simple backend.Server interface.
type MagicServer struct {
	Client *rpc.Client
}

var _ backend.Server = (*MagicServer)(nil)

func NewMagicServer(addr string) backend.Server {
	srv := MagicServer{}
	var err error
	srv.Client, err = rpc.Dial("tcp", addr)
	if err != nil {
		panic(fmt.Sprintf("dialing: %s", err))
	}
	return srv
}

func (srv MagicServer) NewGame(args backend.NewGameRequest) (reply backend.Battlefield) {
	if err := srv.Client.Call("Server.NewGame", args, &reply); err != nil {
		panic(fmt.Sprintf("NewGame: %s", err))
	}
	return reply
}

func (srv MagicServer) BattlefieldState() (reply backend.Battlefield) {
	if err := srv.Client.Call("Server.BattlefieldState", backend.EmptyRequest{}, &reply); err != nil {
		panic(fmt.Sprintf("BattlefieldState: %s", err))
	}
	return
}
