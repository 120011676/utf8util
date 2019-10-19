package hdUnicode

import "strings"

// The location of the fetched substring of large text.
func LargeIndex(str, substr string) int {
	switch strings.Index(str, substr) {
	case 0:
		return 0
	default:
		pos := 0
		i := 0
		str_len := len(str)
		for i < str_len {
			if private_LargeIndex(str, substr, i) {
				return pos
			}
			if str[i] > 127 {
				i += 2
			}
			i++
			pos++
		}
	}
	return -1
}

func private_LargeIndex(str, substr string, i int) bool {
	if str[i] != substr[0] {
		return false
	}
	if len(substr) > 1 {
		return private_LargeIndex(str, substr[1:], i+1)
	}
	return true
}
