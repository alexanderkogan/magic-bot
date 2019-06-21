package backend

type EmptyRequest struct{}

type RPCServer interface {
	NewGame(NewGameRequest, *Battlefield) error
	BattlefieldState(EmptyRequest, *Battlefield) error
}
