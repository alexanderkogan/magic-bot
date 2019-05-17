package main

import "strconv"

func addLifeTotals(you, enemy int, lines []string) []string {
	if len(lines) < 2 {
		return lines
	}
	lines[0] = addTotalOnLine(enemy, lines[0])
	lines[len(lines)-1] = addTotalOnLine(you, lines[len(lines)-1])
	return lines
}

func addTotalOnLine(total int, line string) string {
	if total <= 0 {
		return handleDeadPlayer(line)
	}
	totalString := strconv.Itoa(total)
	if len(line) <= len(totalString) {
		return line
	}
	indent := line[:len(line)-len(totalString)-1]
	return indent + totalString + line[len(line)-1:]
}

const deadPlayer string = "☠"

func handleDeadPlayer(line string) string {
	// Since the deadPlayer placeholder has len 4, it needs to be handled differently.
	if len(line) <= 1 {
		return line
	}
	if len(line) == 2 {
		return deadPlayer + line[1:]
	}
	return line[:len(line)-2] + deadPlayer + line[len(line)-1:]
}
