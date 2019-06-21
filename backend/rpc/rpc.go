package rpc

import (
	"github.com/alexanderkogan/magic-bot/backend"
)

type Server struct {
	Implementation backend.Server
}

var _ backend.RPCServer = (*Server)(nil)

func NewServer() backend.RPCServer {
	return &Server{Implementation: &backend.MockServer{}}
}

func (srv *Server) NewGame(request backend.NewGameRequest, response *backend.Battlefield) error {
	battlefield := srv.Implementation.NewGame(request)
	*response = battlefield
	return nil
}

func (srv Server) BattlefieldState(_ backend.EmptyRequest, response *backend.Battlefield) error {
	battlefield := srv.Implementation.BattlefieldState()
	*response = battlefield
	return nil
}
