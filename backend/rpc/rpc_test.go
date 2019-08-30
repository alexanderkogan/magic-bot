package rpc

import (
	"testing"

	"github.com/alexanderkogan/magic-bot/backend"
)

const defaultYouName = "Liliana Vess"

var testServer = NewServer()

func TestServer_NewGame(t *testing.T) {
	t.Run("starts new game", func(t *testing.T) {
		bf := backend.Battlefield{}
		if err := testServer.NewGame(backend.NewGameRequest{}, &bf); err != nil {
			t.Fatal(err)
		}
		if bf.You.Name != defaultYouName {
			t.Errorf("Expected new game to be started, but got battlefield: %#v", bf)
		}
	})
}

func TestServer_GameStarted(t *testing.T) {
	t.Run("returns false if game not started", func(t *testing.T) {
		testServer = NewServer()
		started := true
		if err := testServer.GameStarted(backend.EmptyRequest{}, &started); err != nil {
			t.Fatal(err)
		}
		if started {
			t.Error("Expected game to not be started, but GameStarted returned true.")
		}
	})

	t.Run("returns true if game started", func(t *testing.T) {
		bf := backend.Battlefield{}
		if err := testServer.NewGame(backend.NewGameRequest{}, &bf); err != nil {
			t.Fatal(err)
		}
		started := false
		if err := testServer.GameStarted(backend.EmptyRequest{}, &started); err != nil {
			t.Fatal(err)
		}
		if !started {
			t.Error("Expected game to be started, but GameStarted returned false.")
		}
	})
}

func TestServer_BattlefieldState(t *testing.T) {
	t.Run("returns new game state", func(t *testing.T) {
		bf := backend.Battlefield{}
		discardBF := backend.Battlefield{}
		if err := testServer.NewGame(backend.NewGameRequest{}, &discardBF); err != nil {
			t.Fatal(err)
		}
		if err := testServer.BattlefieldState(struct{}{}, &bf); err != nil {
			t.Fatal(err)
		}
		if bf.You.Name != defaultYouName {
			t.Errorf("Expected battlefield of started game, but got: %#v", bf)
		}
	})
}
