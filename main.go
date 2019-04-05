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

	fmt.Println(strings.Join(bf, "\n"))
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
	return []string{
		strings.Repeat("-", width),
		"",
		strings.Repeat("-", width),
	}
}
