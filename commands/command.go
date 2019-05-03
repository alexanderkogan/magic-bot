package commands

import "github.com/gdamore/tcell"

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
	Nothing Command = iota
	NewGame
	Quit
)

func Handle(key tcell.EventKey) Command {
	switch key.Key() {
	case tcell.KeyRune:
		switch key.Rune() {
		case 'q':
			return Quit
		case 'n':
			return NewGame
		}
	case tcell.KeyEscape, tcell.KeyEnter:
		return Quit
	}

	return Nothing
}
