package main

import (
	"fmt"
	"strings"

	"github.com/alexanderkogan/magic-bot/backend"
	"github.com/alexanderkogan/magic-bot/backend/rpc"
	"github.com/alexanderkogan/magic-bot/battlefield"
	"github.com/alexanderkogan/magic-bot/commands"
	"github.com/alexanderkogan/magic-bot/tui"
	"github.com/gdamore/tcell"
)

func main() {
	quitMessage := make(chan struct{})
	srv := rpc.NewMagicServer("localhost:42586")
	e := tui.Screen(
		50,
		func(key tcell.EventKey) {
			switch commands.Handle(key) {
			case commands.NewGame:
				srv.NewGame(backend.NewGameRequest{})
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
		lines := getLines(srv, width, height)
		drawScreen(s, lines)
	}
}

func getLines(srv backend.Server, width, height int) []string {
	field := srv.BattlefieldState()
	coms := commands.Commands(width)
	battlefieldLines := battlefield.Battlefield(width, height-len(coms))
	if srv.GameStarted() {
		battlefieldLines = addPlayerNames(field.You.Name, field.Enemy.Name, battlefieldLines)
		battlefieldLines = addLifeTotals(field.You.LifeTotal, field.Enemy.LifeTotal, battlefieldLines)
	}
	battlefieldLines[len(battlefieldLines)/2] = newGameAlertWithIndent(field, width)

	return append(
		battlefieldLines,
		coms...,
	)
}

// TODO this is a hack for having a temporary trigger and will be removed
func newGameAlertWithIndent(field backend.Battlefield, width int) string {
	if field.You.Name != "" {
		newGameMsg := "New Game started"
		indent := width/2 - len(newGameMsg)/2
		if indent < 0 {
			indent = 0
		}
		return strings.Repeat(" ", indent) + newGameMsg
	}
	return ""
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
