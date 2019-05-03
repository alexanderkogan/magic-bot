package commands

import (
	"fmt"
	"strings"
	"testing"

	"github.com/gdamore/tcell"
)

func ExampleCommandsWrap() {
	width := 30
	cm := Commands(width)
	printToConsole(cm)

	// Output: n: New Game - l: Lifepoints
	// q: Quit
}

func ExampleCommandsShouldWrapBestEffort() {
	width := 1
	cm := Commands(width)
	printToConsole(cm)

	// Output: n: New Game
	// l: Lifepoints
	// q: Quit
}

func printToConsole(output []string) {
	fmt.Print(strings.Join(output, "\n"))
}

func TestHandler(t *testing.T) {
	t.Run("should quit on q", func(t *testing.T) {
		got := Handle(*tcell.NewEventKey(tcell.KeyRune, 'q', tcell.ModNone))
		if got != Quit {
			t.Fatalf("Expected to get a quit command (%d), but got %d.", Quit, got)
		}
	})
}
