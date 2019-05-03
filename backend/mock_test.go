package backend

import "testing"

func TestNewGame(t *testing.T) {
	t.Run("empty init", func(t *testing.T) {
		srv := MockServer{}
		battlefield := srv.NewGame(NewGameRequest{})
		expect := Battlefield{
			You:   Player{Name: "Liliana Vess", LifeTotal: 20},
			Enemy: Player{Name: "Chandra Nalaar", LifeTotal: 20},
		}
		if battlefield != expect {
			t.Fatalf("Expected the battlefield to be the default one, but got\n%#v", battlefield)
		}
	})

	t.Run("default battlefield unchanged", func(t *testing.T) {
		srv := MockServer{}
		battlefield := srv.NewGame(NewGameRequest{})
		battlefield.Enemy.LifeTotal = battlefield.Enemy.LifeTotal - 10
		if defaultBattlefield.Enemy.LifeTotal == battlefield.Enemy.LifeTotal {
			t.Fatalf("Changing the battlefield directly should not change internal battlefield, but it did.")
		}
	})

	t.Run("take names from request", func(t *testing.T) {
		srv := MockServer{}
		you := Player{Name: "Jace"}
		enemy := Player{Name: "Garruk"}
		battlefield := srv.NewGame(NewGameRequest{You: you, Enemy: enemy})
		if battlefield.You.Name != you.Name || battlefield.Enemy.Name != enemy.Name {
			t.Fatalf("Expected the names to be set to\nYou: %s\nEnemy: %s\nbut got\nYou: %s\nEnemy: %s",
				you.Name, enemy.Name, battlefield.You.Name, battlefield.Enemy.Name)
		}
	})

	t.Run("take life totals from request", func(t *testing.T) {
		srv := MockServer{}
		you := Player{LifeTotal: 42}
		enemy := Player{LifeTotal: 10}
		battlefield := srv.NewGame(NewGameRequest{You: you, Enemy: enemy})
		if battlefield.You.LifeTotal != you.LifeTotal || battlefield.Enemy.LifeTotal != enemy.LifeTotal {
			t.Fatalf("Expected the life totals to be set to\nYou: %d\nEnemy: %d\nbut got\nYou: %d\nEnemy: %d",
				you.LifeTotal, enemy.LifeTotal, battlefield.You.LifeTotal, battlefield.Enemy.LifeTotal)
		}
	})

	t.Run("sets life totals to default", func(t *testing.T) {
		srv := MockServer{}
		you := Player{Name: "Gideon"}
		enemy := Player{Name: "Ugin"}
		expect := 20
		battlefield := srv.NewGame(NewGameRequest{You: you, Enemy: enemy})
		if battlefield.You.LifeTotal != expect || battlefield.Enemy.LifeTotal != expect {
			t.Fatalf("Expected the life totals to be set to %d but got\nYou: %d\nEnemy: %d",
				expect, battlefield.You.LifeTotal, battlefield.Enemy.LifeTotal)
		}
	})

	t.Run("sets names to default", func(t *testing.T) {
		srv := MockServer{}
		you := Player{LifeTotal: 42}
		enemy := Player{LifeTotal: 10}
		battlefield := srv.NewGame(NewGameRequest{You: you, Enemy: enemy})
		if battlefield.You.Name != defaultBattlefield.You.Name || battlefield.Enemy.Name != defaultBattlefield.Enemy.Name {
			t.Fatalf("Expected the names to be set to\nYou: %s\nEnemy: %s\nbut got\nYou: %s\nEnemy: %s",
				defaultBattlefield.You.Name, defaultBattlefield.Enemy.Name, battlefield.You.Name, battlefield.Enemy.Name)
		}
	})

	t.Run("does not allow negative life totals", func(t *testing.T) {
		srv := MockServer{}
		battlefield := srv.NewGame(NewGameRequest{You: Player{LifeTotal: -2}})
		if battlefield.You.LifeTotal != 20 {
			t.Fatalf("Expected the life total to not be set to a negative number from init, but got %d.", battlefield.You.LifeTotal)
		}
	})
}

func TestBattlefieldState(t *testing.T) {
	t.Run("keep created battlefield", func(t *testing.T) {
		srv := MockServer{}
		you := Player{Name: "Nahiri", LifeTotal: 21}
		enemy := Player{Name: "Sorin", LifeTotal: 22}
		expect := srv.NewGame(NewGameRequest{You: you, Enemy: enemy})
		battlefield := srv.BattlefieldState()
		if expect != battlefield {
			t.Fatalf("Expected the battlefield to stay\n%#v\nbut got\n%#v", expect, battlefield)
		}
	})
}
