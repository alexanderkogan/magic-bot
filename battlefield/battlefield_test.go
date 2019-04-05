package battlefield

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

func printToConsole(output []string) {
	fmt.Print(strings.Join(output, "\n"))
}
