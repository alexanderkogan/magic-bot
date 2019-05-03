package main

import (
	"math"
	"strconv"
	"testing"

	"github.com/gdamore/tcell"
)

func withTestScreen(t *testing.T, test func(tcell.SimulationScreen)) {
	s := tcell.NewSimulationScreen("")
	if s == nil {
		t.Fatalf("Failed to get simulation screen")
	}
	defer s.Fini()
	if e := s.Init(); e != nil {
		t.Fatalf("Failed to initialize screen: %v", e)
	}
	test(s)
}

func TestMainLoop(t *testing.T) {
	t.Run("first screen", func(t *testing.T) {
		withTestScreen(t, func(s tcell.SimulationScreen) {
			mainLoop(s)

			s.Show()
			screenContent, width, height := s.GetContents()
			for position1D, cell := range screenContent {
				x, y := position1DTo2D(position1D, width)
				requireOneRune(t, cell.Runes, x, y)
				checkUpperBorder(t, x, y, cell.Runes[0])
				checkLowerBorder(t, x, y, height, cell.Runes[0])
				checkCommandLine(t, x, y, height, cell.Runes[0])
			}
		})
	})
}

func requireOneRune(t *testing.T, runes []rune, x, y int) {
	if len(runes) > 1 {
		t.Fatalf("Unexpected number of runes in %d, %d: %v", x, y, runes)
	}
}

func checkUpperBorder(t *testing.T, x, y int, content rune) {
	if y == 0 {
		if content != '-' {
			t.Fatalf("Expected upper border to be filled with '-', but got %s at (%d, %d)", strconv.QuoteRune(content), x, y)
		}
	}
}

func checkLowerBorder(t *testing.T, x, y, height int, content rune) {
	if y == height-2 {
		if content != '-' {
			t.Fatalf("Expected lower border of battlefield to be filled with '-', but got %s at (%d, %d)", strconv.QuoteRune(content), x, y)
		}
	}
}

func checkCommandLine(t *testing.T, x, y, height int, content rune) {
	var commandLineSnapshot = []rune("n: New Game - l: Lifepoints - q: Quit")

	if y == height-1 {
		restOfLine := x >= len(commandLineSnapshot)
		if !restOfLine && content != rune(commandLineSnapshot[x]) {
			t.Fatalf("Expected command line to be '%s' but got '%s' instead of '%s' at (%d, %d).", string(commandLineSnapshot), string(content), string(commandLineSnapshot[x]), x, y)
		}
		if restOfLine && content != ' ' {
			t.Fatalf("Expected rest of command line to be empty, but got '%s' at (%d, %d).", string(content), x, y)
		}
	}
}

func position1DTo2D(pos, width int) (x int, y int) {
	x = pos % width
	y = int(math.Floor(float64(pos) / float64(width)))
	return
}
