package main

func addPlayerNames(youName, enemyName string, lines []string) []string {
	if len(lines) < 2 {
		return lines
	}
	lines[0] = addNameOnLine(enemyName, lines[0])
	lines[len(lines)-1] = addNameOnLine(youName, lines[len(lines)-1])
	return lines
}

func addNameOnLine(name string, line string) string {
	if name != "" && len(line) > 0 {
		restOfLine := ""
		nameFitsOnLine := 1+len(name) < len(line)
		if nameFitsOnLine {
			restOfLine = line[1+len(name):]
		}
		nameToPrint := name
		if len(name)+1 > len(line) {
			nameToPrint = name[:len(line)-1]
		}
		return line[:1] + nameToPrint + restOfLine
	}
	return line
}
