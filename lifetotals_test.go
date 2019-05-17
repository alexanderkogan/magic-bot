package main

import (
	"strconv"
	"testing"
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
		expect := []string{"------" + strconv.Itoa(enemy) + "-", line2, ",,,,,,," + strconv.Itoa(you) + ","}
		got := addLifeTotals(you, enemy, []string{line1, line2, line3})
		if len(got) != 3 || got[0] != expect[0] || got[1] != expect[1] || got[2] != expect[2] {
			t.Fatalf("Expected life totals to be added to the right of the line "+
				"and rest of lines to be untouched, but got %#v\nExpected: %#v", got, expect)
		}
	})
}
