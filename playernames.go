package main

import "github.com/alexanderkogan/magic-bot/backend"

func addPlayerNames(field backend.Battlefield, lines []string) []string {
	lines[0] = addNameOnLine(field.Enemy.Name, lines[0])
	lines[len(lines)-1] = addNameOnLine(field.You.Name, lines[len(lines)-1])
	return lines
}

func addNameOnLine(name string, line string) string {
	if name != "" {
		restOfLine := ""
		nameFitsOnLine := 1+len(name) < len(line)
		if nameFitsOnLine {
			restOfLine = line[1+len(name):]
		}
		return line[:1] + name + restOfLine
	}
	return line
}
