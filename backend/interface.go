package backend

type NewGameRequest struct {
	You   Player
	Enemy Player
}

type Battlefield struct {
	You   Player
	Enemy Player
}

type Player struct {
	Name      string
	LifeTotal int
	Library   DeckID
}

type DeckID string

type Server interface {
	NewGame(NewGameRequest) Battlefield

	GameStarted() bool

	BattlefieldState() Battlefield
}
