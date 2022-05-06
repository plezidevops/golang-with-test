package iteration

import "strings"

const repeatCount = 5

// Repeat a character n number of times
func Repeat(text string) string {
	return strings.Repeat(text, repeatCount)
}
