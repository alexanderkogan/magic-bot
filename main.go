package main

import (
	"fmt"
	"strings"

	"github.com/alexanderkogan/magic-bot/tui"
	"github.com/gdamore/tcell"
)

func main() {
	quitMessage := make(chan struct{})
	e := tui.Screen(
		50,
		func(key tcell.EventKey) {
			switch handleCommands(key) {
			case Quit:
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
	bf := Battlefield(width)
	commands := Commands(width)
	fmt.Println(strings.Join(append(bf, commands...), "\n"))
}

func Battlefield(width int) []string {
	border := strings.Repeat("-", width)
	return []string{
		border,
		"",
		border,
	}
}

func Commands(width int) []string {
	cmds := []string{
		"n: New Game",
		"l: Lifepoints",
		"q: Quit",
	}
	lines := toLinesByWords(cmds, width)
	return lines
}

func toLinesByWords(words []string, lineWidth int) []string {
	var lines []string
	var line string
	cmdDivider := " - "
	for _, value := range words {
		if len(line)+len(cmdDivider)+len(value) > lineWidth {
			lines = append(lines, line)
			line = value
			continue
		}
		if line == "" {
			line = value
		} else {
			line = line + cmdDivider + value
		}
	}
	if line != "" {
		lines = append(lines, line)
	}
	return lines
}

type Command int

const (
	Quit Command = iota
	Nothing
)

func handleCommands(key tcell.EventKey) Command {
	switch key.Key() {
	case tcell.KeyRune:
		switch key.Rune() {
		case 'q':
			return Quit
		}
	case tcell.KeyEscape, tcell.KeyEnter:
		return Quit
	}

	return Nothing
}
