package main

import (
	"strconv"
	"testing"

	"github.com/alexanderkogan/magic-bot/backend"
	"github.com/gdamore/tcell"
)

func TestAddLifeTotals(t *testing.T) {
	t.Run("nil line slice", func(t *testing.T) {
		got := addLifeTotals(2, 2, nil)
		if got != nil {
			t.Fatalf("Expected addLifeTotals to do nothing, if nil lines provided, but got %#v", got)
		}
	})

	t.Run("small line slice", func(t *testing.T) {
		got := addLifeTotals(2, 2, []string{""})
		if len(got) != 1 || got[0] != "" {
			t.Fatalf("Expected addLifeTotals to do nothing, if not enough lines provided, but got %#v", got)
		}
	})

	t.Run("line too short", func(t *testing.T) {
		got := addLifeTotals(2, 20, []string{"--", "-"})
		if len(got) != 2 || got[0] != "--" || got[1] != "-" {
			t.Fatalf("Expected addLifeTotals to do nothing, if lines are too short, but got %#v", got)
		}
	})

	t.Run("dead player line too short", func(t *testing.T) {
		got := addLifeTotals(-2, 0, []string{"-", "-"})
		if len(got) != 2 || got[0] != "-" || got[1] != "-" {
			t.Fatalf("Expected addLifeTotals to do nothing, if lines are too short for dead players, but got %#v", got)
		}
	})

	t.Run("dead player", func(t *testing.T) {
		got := addLifeTotals(0, -1, []string{"---", "----", ".."})
		if got[0] != "-"+deadPlayer+"-" || got[2] != ""+deadPlayer+"." {
			t.Fatalf("Expected totals to show %s but got %#v", deadPlayer, got)
		}
	})

	t.Run("correct placement", func(t *testing.T) {
		line1, line2, line3 := "---", "...", ",,,"
		you, enemy := 20, 42
		expect := []string{strconv.Itoa(enemy) + "-", line2, strconv.Itoa(you) + ","}
		got := addLifeTotals(you, enemy, []string{line1, line2, line3})
		if len(got) != 3 || got[0] != expect[0] || got[1] != expect[1] || got[2] != expect[2] {
			t.Fatalf("Expected your life total to be added on lower line and enemy life total to be added on upper line "+
				"and rest of lines to be untouched, but got %#v\nExpected: %#v", got, expect)
		}
	})

	t.Run("line continuation", func(t *testing.T) {
		line1, line2, line3 := "----------", "..........", ",,,,,,,,,,"
		you, enemy := 20, 120
		expect := []string{"----|" + strconv.Itoa(enemy) + "|-", line2, ",,,,,|" + strconv.Itoa(you) + "|,"}
		got := addLifeTotals(you, enemy, []string{line1, line2, line3})
		if len(got) != 3 || got[0] != expect[0] || got[1] != expect[1] || got[2] != expect[2] {
			t.Fatalf("Expected life totals to be added to the right of the line "+
				"and rest of lines to be untouched, but got %#v\nExpected: %#v", got, expect)
		}
	})
}

func TestLifeTotalsOnScreen(t *testing.T) {
	t.Run("should show life totals", func(t *testing.T) {
		withTestScreen(t, func(screen tcell.SimulationScreen) {
			srv := &backend.MockServer{}
			you, enemy := backend.Player{Name: "Alex", LifeTotal: 9}, backend.Player{Name: "Niko", LifeTotal: 200}
			srv.NewGame(backend.NewGameRequest{You: you, Enemy: enemy})
			mainLoop(srv)(screen)

			screenContent, width, height := screen.GetContents()
			youHeight, enemyHeight := height-2, 0
			for position1D, cell := range screenContent {
				x, y := position1DTo2D(position1D, width)
				requireOneRune(t, cell.Runes, x, y)
				checkLifeTotal(t, x, y, surroundLifeTotalWithBrackets(strconv.Itoa(you.LifeTotal)), youHeight, width, cell.Runes[0])
				checkLifeTotal(t, x, y, surroundLifeTotalWithBrackets(strconv.Itoa(enemy.LifeTotal)), enemyHeight, width, cell.Runes[0])
			}
		})
	})

	t.Run("should show dead player", func(t *testing.T) {
		withTestScreen(t, func(screen tcell.SimulationScreen) {
			srv := &backend.MockServer{}
			you, enemy := backend.Player{Name: "Alex"}, backend.Player{Name: "Niko"}
			srv.NewGame(backend.NewGameRequest{You: you, Enemy: enemy})
			srv.OverwritePlayerLifeTotal(0, -1)
			mainLoop(srv)(screen)

			screenContent, width, height := screen.GetContents()
			youHeight, enemyHeight := height-2, 0
			for position1D, cell := range screenContent {
				x, y := position1DTo2D(position1D, width)
				requireOneRune(t, cell.Runes, x, y)
				checkLifeTotal(t, x, y, deadPlayer, youHeight, width, cell.Runes[0])
				checkLifeTotal(t, x, y, deadPlayer, enemyHeight, width, cell.Runes[0])
			}
		})
	})

	// TODO check overwriting of name
	// TODO fix bug, that lifetotal not showing on narrow screens
}

func checkLifeTotal(t *testing.T, x, y int, lifeTotal string, expectedHeight, width int, content rune) {
	lifeTotalLen := getLifeLen(lifeTotal)
	startX, endX := getStartEndForLife(width, lifeTotalLen)
	if y == expectedHeight && x >= startX-2 {
		placeOfLife := x >= startX && x <= endX
		beforeLife := x >= startX-2 && x < startX
		afterLife := x > endX+1

		if placeOfLife {
			livingPlayerOK := lifeTotal == deadPlayer || content == rune(lifeTotal[x-startX])
			deadPlayerOK := lifeTotal != deadPlayer || content == deadPlayerRune

			if !livingPlayerOK || !deadPlayerOK {
				t.Errorf("Expected '%s' to be printed here, but got '%s' at (%d, %d).", lifeTotal, string(content), x, y)
			}
		}
		if (beforeLife || afterLife) && content != '-' {
			t.Fatalf("Expected line around the life total to be filled with '-' but got '%s' at (%d, %d).", string(content), x, y)
		}
	}
}

func getLifeLen(life string) int {
	if life == deadPlayer {
		return 1
	}
	return len(life)
}

func getStartEndForLife(width, lifeTotalLen int) (int, int) {
	start := width - 1 - lifeTotalLen
	end := start + lifeTotalLen - 1
	return start, end
}
