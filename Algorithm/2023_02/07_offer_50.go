package main

func main() {

}
func firstUniqChar(s string) byte {
	var searchLocation = make(map[byte]int)
	var searchBytes = make(map[int]byte)
	var index = 0
	var location = make([]int, len(s))
	for i := 0; i < len(s); i++ {
		if value, ok := searchLocation[s[i]]; ok {
			location[value]++
			continue
		}
		searchLocation[s[i]] = index
		searchBytes[index] = s[i]
		location[index]++
		index++
	}
	for i := 0; i < len(location); i++ {
		if location[i] == 1 {
			return searchBytes[i]
		}
	}
	return ' '

}
