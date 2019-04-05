package main

import (
	"fmt"
	"strings"

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
	width, _ := s.Size()
	bf := battlefield.Battlefield(width)
	coms := commands.Commands(width)
	lines := append(bf, coms...)
	fmt.Print(strings.Join(lines, "\r\n"))
}
