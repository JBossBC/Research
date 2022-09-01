package main

import "strings"

func main() {
	print(isPrefixOfWord("i love eating burger", "burg"))
}
func isPrefixOfWord(sentence string, searchWord string) int {
	var sentenceArr = strings.Split(sentence, " ")
	for i := 0; i < len(sentenceArr); i++ {
		var success = true
		var sentenceValue = sentenceArr[i]
		if len(sentenceValue) < len(searchWord) {
			continue
		}
		for j := 0; j < len(sentenceArr[i]); j++ {
			if j < len(searchWord) && sentenceValue[j] != searchWord[j] {
				success = false
				break
			}
		}
		if success {
			return i + 1
		}

	}
	return -1
}
