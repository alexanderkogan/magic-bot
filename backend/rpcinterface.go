package backend

type EmptyRequest struct{}

// TODO Rename functions to have less confusion?
type RPCServer interface {
	NewGame(NewGameRequest, *Battlefield) error
	GameStarted(EmptyRequest, *bool) error
	BattlefieldState(EmptyRequest, *Battlefield) error
}
