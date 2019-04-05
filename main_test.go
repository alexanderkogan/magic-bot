package main

import (
	"fmt"
	"strings"
)

func ExampleEmptyBattlefield() {
	width := 10
	bf := Battlefield(width)
	fmt.Print(strings.Join(bf, "\n"))

	// Output: ----------
	//
	// ----------
}
