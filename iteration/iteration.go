package iteration

import "strings"

func Repeat(character string) string {
	var repeated strings.Builder
	for i := 1; i <= 5; i++ {
		repeated.WriteString(character)
	}
	return repeated.String()
}
