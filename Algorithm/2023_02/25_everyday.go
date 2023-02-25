package main

func main() {
	minimumSwap("xx", "yy")
}
func minimumSwap(s1 string, s2 string) int {
	var diff = make([]byte, len(s2))
	var mapping = map[byte]byte{'x': 'y', 'y': 'x'}
	for i := 0; i < len(s2); i++ {
		if s1[i] != s2[i] {
			diff[i] = mapping[s2[i]]
		}
	}
	var times int
	var ignore = make(map[int]bool)
	for i := 0; i < len(s2); i++ {
		if s1[i] != s2[i] {
			if ignore[i] {
				continue
			}
			var find = false
			for j := 0; j < len(diff); j++ {
				if diff[j] == ' ' {
					continue
				}
				if diff[j] == s1[i] {
					find = true
					diff[j] = diff[i]
					diff[i] = ' '
					if diff[j] == s1[j] {
						ignore[j] = true
						diff[j] = ' '
					}
					times++
					break
				}
			}
			if !find {
				return -1
			}
		}
	}
	return times
}
