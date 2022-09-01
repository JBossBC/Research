package main

/**
leetcode 1422
*/
func main() {
	println(maxScore("00"))
}

func maxScore(s string) int {
	var maxScore = 0
	var tempScore = 0
	if len(s) <= 1 {
		return 0
	}
	if len(s) == 2 {
		var result = 0
		if s[0] == '0' {
			result++
		}
		if s[1] == '1' {
			result++
		}
		return result
	}
	for i := 1; i < len(s)-1; i++ {
		tempScore = 0
		for j := i; j >= 0; j-- {
			if s[j] == '0' {
				tempScore++
			}
		}
		for j := i; j < len(s); j++ {
			if s[j] == '1' {
				tempScore++
			}
		}
		if tempScore > maxScore {
			maxScore = tempScore
		}
	}
	return maxScore
}
func dynamic(s string) {

}
