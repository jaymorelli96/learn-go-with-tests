package iteration

func Repeat(word string, n int) string {
	var result string
	for i := 0; i < n; i++ {
		result += word
	}

	return result
}
