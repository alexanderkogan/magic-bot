package main

import (
	"fmt"

	"github.com/alexanderkogan/magic-bot/battlefield"
	"github.com/alexanderkogan/magic-bot/commands"
	"github.com/alexanderkogan/magic-bot/tui"
	"github.com/gdamore/tcell"
)

func main() {
	quitMessage := make(chan struct{})
	e := tui.Screen(
		50,
		func(key tcell.EventKey) {
			switch commands.Handle(key) {
			case commands.Quit:
				close(quitMessage)
			}
		},
		mainLoop,
		quitMessage,
	)
	if e != nil {
		fmt.Println(e)
	}
}

func mainLoop(s tcell.Screen) {
	s.Sync()
	lines := getLines(s)
	drawScreen(s, lines)
}

func getLines(s tcell.Screen) []string {
	width, height := s.Size()
	coms := commands.Commands(width)
	return append(
		battlefield.Battlefield(width, height-len(coms)),
		coms...,
	)
}

func drawScreen(s tcell.Screen, lines []string) {
	var noCombiningRunes []rune

	defer s.Show()
	s.Clear()

	for y, lineToDraw := range lines {
		for x, characterToDraw := range lineToDraw {
			s.SetContent(x, y, characterToDraw, noCombiningRunes, tcell.StyleDefault)
		}
	}
}
