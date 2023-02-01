package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(decodeMessage("the quick brown fox jumps over the lazy dog", "vkbs bs t suepuv"))
}

var alph = []rune{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z'}

func decodeMessage(key string, message string) string {
	var memory = make(map[rune]rune)
	var newStr = strings.Builder{}
	//make dict
	func() {
		var diffDict = make(map[byte]any, len(alph))
		for i := 0; i < len(key); i++ {
			if key[i] == ' ' {
				continue
			}
			if _, ok := diffDict[key[i]]; !ok {
				newStr.WriteRune(rune(key[i]))
				diffDict[key[i]] = nil
				if newStr.Len() > len(alph) {
					return
				}
			}
		}
	}()
	var newString = newStr.String()
	var changeFunc = func(s rune) rune {
		if s == ' ' {
			return s
		}
		if value, ok := memory[s]; ok {
			return value
		}
		for i := 0; i < len(newString); i++ {
			if rune(newString[i]) == s {
				var result = alph[i%26]
				memory[s] = result
				return result
			}
		}
		return ' '
	}
	return strings.Map(changeFunc, message)

}
