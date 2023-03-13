package iteration

import "strings"

func Repeat(word string, n int) string {
	var result strings.Builder
	for i := 0; i < n; i++ {
		result.Write([]byte(word))
	}

	return result.String()
}
