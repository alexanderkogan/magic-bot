package backend

type NewGameRequest struct {
	You   Player
	Enemy Player
}

type Battlefield struct {
	You   Player
	Enemy Player
	Zone  Zones
}
type Zones struct {
	Battlefield BattlefieldZone
}

type Player struct {
	Name      string
	LifeTotal int
	Library   DeckID
}

type DeckID string

type BattlefieldZone struct {
	You   []Card
	Enemy []Card
}

type Card struct {
	ID     CardID
	Tapped bool
}

type CardID struct {
	Set              SetID
	CollectorsNumber int
}

type SetID string

// Standard until Autumn 2019
const (
	Ixalan                              SetID = "XLN"
	RivalsOfIxalan                            = "RIX"
	Dominaria                                 = "DOM"
	GlobalSeriesJiangYangguAndMuYanling       = "GS1"
	CoreSet2019                               = "M19"
)

// Standard until Autumn 2020
const (
	GuildsOfRavnica   SetID = "GRN"
	RavnicaAllegiance       = "RNA"
	WarOfTheSpark           = "WAR"
	CoreSet2020             = "M20"
)

type Server interface {
	NewGame(NewGameRequest) Battlefield

	GameStarted() bool

	BattlefieldState() Battlefield
}
