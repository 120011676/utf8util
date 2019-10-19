package hdUnicode

import "strings"

// LastIndex returns the index of the last instance of substr in s, or -1 if substr is not present in s.
func LastIndex(str, substr string) int {
	switch strings.LastIndex(str, substr) {
	case 0:
		return 0
	default:
		str_len := len(str)
		substr_len := len(substr)
		pos := len([]rune(str)) - len([]rune(substr))
		i := str_len - 1
		for i > 0 {
			if str[i] == substr[substr_len-1] {
				if str[i-1] == substr[substr_len-2] {
					if str[i-2] == substr[substr_len-3] {
						if str[i-substr_len+1:i+1] == substr {
							return pos
						}
					}
				}
			}
			if str[i] > 127 {
				i -= 2
			}
			i--
			pos--
		}
	}
	return -1
}
