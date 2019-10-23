package utf8Util

import "strings"

// LastIndexOne returns the index of the last instance of substr in s, or -1 if substr is not present in s.
func LastIndexOne(str, substr string) int {
	switch strings.LastIndex(str, substr) {
	case 0:
		return 0
	default:
		str_len := len(str)
		pos := len([]rune(str)) - len([]rune(substr))
		i := str_len - 1
		for i > 0 {
			if str[i] == substr[2] {
				if str[i-1] == substr[1] {
					if str[i-2] == substr[0] {
						return pos
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
