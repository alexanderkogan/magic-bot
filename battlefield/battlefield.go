package battlefield

import "strings"

func Battlefield(width int) []string {
	border := strings.Repeat("-", width)
	return []string{
		border,
		"",
		border,
	}
}
