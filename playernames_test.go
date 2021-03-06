package main

import (
	"testing"

	"github.com/alexanderkogan/magic-bot/backend"
	"github.com/gdamore/tcell"
)

func TestAddNameOnLine(t *testing.T) {
	t.Run("should cut off name if line too short", func(t *testing.T) {
		line := "----------"
		got := addNameOnLine("absurdly long name", line)
		if len(got) != len(line) {
			t.Fatalf("Expected addNameOnLine to cut off the name after a length of %d, but got %#v (length: %d)", len(line), got, len(got))
		}
	})
}

func TestAddPlayerNames(t *testing.T) {
	t.Run("nil line slice", func(t *testing.T) {
		got := addPlayerNames("", "", nil)
		if got != nil {
			t.Fatalf("Expected addPlayerNames to do nothing, if nil lines provided, but got %#v", got)
		}
	})

	t.Run("empty you name", func(t *testing.T) {
		lines := []string{"-----", "-----", "-----"}
		got := addPlayerNames("", "enemy", lines)
		expect := []string{"-enem", lines[1], lines[2]}
		if len(got) != 3 || got[0] != expect[0] || got[1] != expect[1] || got[2] != expect[2] {
			t.Fatalf("Expected addPlayerNames to add no you name, if empty name is provided, but got %#v\nExpected: %#v", got, expect)
		}
	})

	t.Run("empty enemy name", func(t *testing.T) {
		lines := []string{"-----", "-----", "-----"}
		got := addPlayerNames("you", "", lines)
		expect := []string{lines[0], lines[1], "-you-"}
		if len(got) != 3 || got[0] != expect[0] || got[1] != expect[1] || got[2] != expect[2] {
			t.Fatalf("Expected addPlayerNames to add no you name, if empty name is provided, but got %#v\nExpected: %#v", got, expect)
		}
	})

	t.Run("small line slice", func(t *testing.T) {
		got := addPlayerNames("", "", []string{""})
		if len(got) != 1 || got[0] != "" {
			t.Fatalf("Expected addPlayerNames to do nothing, if not enough lines provided, but got %#v", got)
		}
	})

	t.Run("line too short", func(t *testing.T) {
		got := addPlayerNames("Player 1", "Player 2", []string{"", ""})
		if len(got) != 2 || got[0] != "" || got[1] != "" {
			t.Fatalf("Expected addPlayerNames to do nothing, if lines are too short, but got %#v", got)
		}
	})

	t.Run("correct placement", func(t *testing.T) {
		line1, line2, line3 := "------", "....", ",,,,"
		you, enemy := "You", "Enemy"
		expect := []string{"-Enemy", line2, ",You"}
		got := addPlayerNames(you, enemy, []string{line1, line2, line3})
		if len(got) != 3 || got[0] != expect[0] || got[1] != expect[1] || got[2] != expect[2] {
			t.Fatalf("Expected your name to be added on lower line and enemy name to be added on upper line "+
				"and rest of lines to be untouched, but got %#v\nExpected: %#v", got, expect)
		}
	})

	t.Run("line continuation", func(t *testing.T) {
		line1, line2, line3 := "----------", "..........", ",,,,,,,,,,"
		you, enemy := "You", "Enemy"
		expect := []string{"-" + enemy + "----", line2, "," + you + ",,,,,,"}
		got := addPlayerNames(you, enemy, []string{line1, line2, line3})
		if len(got) != 3 || got[0] != expect[0] || got[1] != expect[1] || got[2] != expect[2] {
			t.Fatalf("Expected lines to be continued around names, but got %#v\nExpected: %#v", got, expect)
		}
	})
}

func TestPlayerNamesOnScreen(t *testing.T) {
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
		withTestScreenOfSize(t, 10, 25, func(screen tcell.SimulationScreen) {
			srv := &backend.MockServer{}
			you := backend.Player{Name: "Your Name", LifeTotal: 999999999999} // The life total is longer than the line, thus not showing up.
			srv.NewGame(backend.NewGameRequest{You: you})
			mainLoop(srv)(screen)

			screenContent, width, height := screen.GetContents()
			youHeight := height - 5 // FIXME This should be one less. See (https://github.com/alexanderkogan/magic-bot/issues/8)
			for position1D, cell := range screenContent {
				x, y := position1DTo2D(position1D, width)
				requireOneRune(t, cell.Runes, x, y)
				checkPlayerName(t, x, y, you.Name, youHeight, cell.Runes[0])
			}
		})
	})
}

func checkPlayerName(t *testing.T, x, y int, name string, expectedHeight int, content rune) {
	startX := 1
	if y == expectedHeight && x >= startX {
		endX := startX + len(name) - 1
		placeOfName := x >= startX && x <= endX
		afterName := x > endX && x <= endX+2
		beforeName := x < startX
		if placeOfName && content != rune(name[x-startX]) {
			t.Errorf("Expected '%s' to be printed here, but got '%s' at (%d, %d).", name, string(content), x, y)
		}
		if (afterName || beforeName) && content != '-' {
			t.Fatalf("Expected line around the name to be filled with '-' but got '%s' at (%d, %d).", string(content), x, y)
		}
	}
}
