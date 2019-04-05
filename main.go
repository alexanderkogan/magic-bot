package main

import (
	"fmt"
	"os"
	"strings"

	"golang.org/x/crypto/ssh/terminal"
)

func main() {
	width := getTerminalWidth()
	bf := Battlefield(width)
	commands := Commands(width)

	fmt.Println(strings.Join(append(bf, commands...), "\n"))
}

func getTerminalWidth() int {
	fd := int(os.Stdout.Fd())
	width, _, err := terminal.GetSize(fd)
	if err != nil {
		panic(err)
	}
	return width
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
	var lines []string
	var line string
	cmdDivider := " - "
	for _, value := range cmds {
		if len(line)+len(cmdDivider)+len(value) > width {
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
