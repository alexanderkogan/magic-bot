package main

import (
	"testing"

	"github.com/alexanderkogan/magic-bot/backend"
	"github.com/gdamore/tcell"
)

func TestNewGamePlayerNames(t *testing.T) {
	t.Run("new game", func(t *testing.T) {
		t.Run("should show player names", func(t *testing.T) {
			withTestScreen(t, func(screen tcell.SimulationScreen) {
				srv := &backend.MockServer{}
				you, enemy := backend.Player{Name: "Alex"}, backend.Player{Name: "Niko"}
				srv.NewGame(backend.NewGameRequest{You: you, Enemy: enemy})
				mainLoop(srv)(screen)

				screenContent, width, height := screen.GetContents()
				youHeight, enemyHeight := height-2, 0
				for position1D, cell := range screenContent {
					x, y := position1DTo2D(position1D, width)
					requireOneRune(t, cell.Runes, x, y)
					checkPlayerName(t, x, y, you.Name, youHeight, cell.Runes[0])
					checkPlayerName(t, x, y, enemy.Name, enemyHeight, cell.Runes[0])
				}
			})
		})
		t.Run("should move player name up, if command takes multiple lines", func(t *testing.T) {
			withTestScreenOfSize(t, 5, 25, func(screen tcell.SimulationScreen) {
				srv := &backend.MockServer{}
				srv.NewGame(backend.NewGameRequest{})
				mainLoop(srv)(screen)

				screenContent, width, height := screen.GetContents()
				you := "Liliana Vess"
				youHeight := height - 5 // FIXME This should be one less. See (https://github.com/alexanderkogan/magic-bot/issues/8)
				for position1D, cell := range screenContent {
					x, y := position1DTo2D(position1D, width)
					requireOneRune(t, cell.Runes, x, y)
					checkPlayerName(t, x, y, you, youHeight, cell.Runes[0])
				}
			})
		})
	})
}

func checkPlayerName(t *testing.T, x, y int, name string, expectedHeight int, content rune) {
	startX := 1
	if y == expectedHeight && x >= startX {
		placeOfName := x >= startX && x < startX+len(name)
		if placeOfName && content != rune(name[x-startX]) {
			t.Errorf("Expected '%s' to be printed here, but got '%s' at (%d, %d).", name, string(content), x, y)
		}
		if !placeOfName && content != '-' {
			t.Fatalf("Expected rest of line to be filled with '-' but got '%s' at (%d, %d).", string(content), x, y)
		}
	}
}
