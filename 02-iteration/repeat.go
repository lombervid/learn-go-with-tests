package iteration

import "strings"

func Repeat(character string, times int) string {
	// return strings.Repeat(character, times)
	var repeated strings.Builder
	for range times {
		repeated.WriteString(character)
	}
	return repeated.String()
}
