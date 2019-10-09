package hdUnicode

func Index(str, substr string) int {
	pos := 0
	i := 0
	substr_len := len(substr)
	str_len := len(str)
	for i < str_len {
		if str[i] == substr[0] {
			if str[i:i+substr_len] == substr {
				return pos
			}
		}
		if str[i] > 127 {
			i += 2
		}
		i++
		pos++
	}
	return -1
}
