package backend

type MockServer struct {
	battlefield Battlefield
}

var _ Server = (*MockServer)(nil)

var defaultBattlefield = Battlefield{
	You:   Player{Name: "Liliana Vess", LifeTotal: 20},
	Enemy: Player{Name: "Chandra Nalaar", LifeTotal: 20},
}

func (srv *MockServer) NewGame(init NewGameRequest) Battlefield {
	srv.battlefield = defaultBattlefield
	if init.You.Name != "" {
		srv.battlefield.You.Name = init.You.Name
	}
	if init.Enemy.Name != "" {
		srv.battlefield.Enemy.Name = init.Enemy.Name
	}
	if init.You.LifeTotal > 0 {
		srv.battlefield.You.LifeTotal = init.You.LifeTotal
	}
	if init.Enemy.LifeTotal > 0 {
		srv.battlefield.Enemy.LifeTotal = init.Enemy.LifeTotal
	}
	return srv.battlefield
}

func (srv MockServer) BattlefieldState() Battlefield {
	return srv.battlefield
}