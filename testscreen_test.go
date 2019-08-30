package main

import (
	"math"
	"testing"

	"github.com/gdamore/tcell"
)

func withTestScreen(t *testing.T, test func(tcell.SimulationScreen)) {
	withTestScreenOfSize(t, 80, 25, test)
}
func withTestScreenOfSize(t *testing.T, width, height int, test func(tcell.SimulationScreen)) {
	s := tcell.NewSimulationScreen("")
	if s == nil {
		t.Fatalf("Failed to get simulation screen")
	}
	defer s.Fini()
	if e := s.Init(); e != nil {
		t.Fatalf("Failed to initialize screen: %v", e)
	}
	s.SetSize(width, height)
	test(s)
}

func position1DTo2D(pos, width int) (x int, y int) {
	x = pos % width
	y = int(math.Floor(float64(pos) / float64(width)))
	return
}

func requireOneRune(t *testing.T, runes []rune, x, y int) {
	if len(runes) > 1 {
		t.Fatalf("Unexpected number of runes in %d, %d: %v", x, y, runes)
	}
}
