package main

func main() {

}
func lengthOfLongestSubstring(s string) int {
	var maxLength = 0
	for i := 0; i < len(s); i++ {
		var memory = make(map[byte]int)
		var length = 0
		for j := i; j < len(s); j++ {
			var temp = s[j]
			if value, ok := memory[temp]; ok {
				i = value
				break
			}
			memory[temp] = j
			length++
		}
		if length > maxLength {
			maxLength = length
		}
	}
	return maxLength
}
