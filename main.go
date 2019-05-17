package main

import (
	"fmt"

	"github.com/alexanderkogan/magic-bot/backend"
	"github.com/alexanderkogan/magic-bot/battlefield"
	"github.com/alexanderkogan/magic-bot/commands"
	"github.com/alexanderkogan/magic-bot/tui"
	"github.com/gdamore/tcell"
)

func main() {
	quitMessage := make(chan struct{})
	srv := &backend.MockServer{}
	e := tui.Screen(
		50,
		func(key tcell.EventKey) {
			switch commands.Handle(key) {
			case commands.Quit:
				close(quitMessage)
			}
		},
		mainLoop(srv),
		quitMessage,
	)
	if e != nil {
		fmt.Println(e)
	}
}

func mainLoop(srv backend.Server) func(tcell.Screen) {
	return func(s tcell.Screen) {
		s.Sync()
		width, height := s.Size()
		lines := getLines(srv.BattlefieldState(), width, height)
		drawScreen(s, lines)
	}
}

func getLines(field backend.Battlefield, width, height int) []string {
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
