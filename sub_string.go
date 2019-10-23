package utf8Util

// Gets the substring, gets it by starting and closing the index, and returns the substring.
func SubString(str string, start, end int) string {
	return string([]rune(str)[start:end])
}
