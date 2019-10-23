package utf8Util

import "strings"

// IndexOne returns the index of the first instance of substr in s, or -1 if substr is not present in s.
func IndexOne(str, substr string) int {
	switch strings.Index(str, substr) {
	case 0:
		return 0
	default:
		pos := 0
		i := 0
		str_len := len(str)
		for i < str_len {
			if str[i] == substr[0] {
				if str[i+1] == substr[1] {
					if str[i+2] == substr[2] {
						return pos
					}
				}
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
