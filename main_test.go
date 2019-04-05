package main

import (
	"fmt"
	"strings"
)

func ExampleEmptyBattlefield() {
	width := 10
	bf := Battlefield(width)
	printToConsole(bf)

	// Output: ----------
	//
	// ----------
}

func ExampleCommandsWrap() {
	width := 30
	cm := Commands(width)
	printToConsole(cm)

	// Output: n: New Game - l: Lifepoints
	// q: Quit
}

func ExampleCommandsShouldWrapBestEffort() {
	width := 1
	cm := Commands(width)
	printToConsole(cm)

	// Output: n: New Game
	// l: Lifepoints
	// q: Quit
}

func printToConsole(output []string) {
	fmt.Print(strings.Join(output, "\n"))
}
