package battlefield

import "strings"

func Battlefield(width, height int) []string {
	border := strings.Repeat("-", width)
	out := []string{border}
	for line := 0; line < height-2; line++ {
		out = append(out, "")
	}
	return append(out, border)
}
