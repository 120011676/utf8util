package hdUnicode

// Gets the substring, gets it by starting the index and the length of the substring, and returns the substring.
func SubStr(str string, index, len int) string {
	return string([]rune(str)[index : index+len-1])
}
