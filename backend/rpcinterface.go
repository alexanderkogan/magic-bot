package backend

type EmptyRequest struct{}

type RPCServer interface {
	NewGame(NewGameRequest, *Battlefield) error
	GameStarted(EmptyRequest, *bool) error
	BattlefieldState(EmptyRequest, *Battlefield) error
}
