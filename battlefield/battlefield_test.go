package battlefield

import (
	"fmt"
	"strings"
)

func ExampleEmptyBattlefield() {
	width, height := 10, 3
	bf := Battlefield(width, height)
	printToConsole(bf)

	// Output: ----------
	//
	// ----------
}

func printToConsole(output []string) {
	fmt.Print(strings.Join(output, "\n"))
}
