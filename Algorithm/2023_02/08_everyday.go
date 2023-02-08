package main

import "fmt"

func main() {

	var result = removeSubfolders2([]string{"/ah/al/am", "/ah/al"})
	for i := 0; i < len(result); i++ {
		fmt.Printf("%s ", result[i])
	}
}

func removeSubfolders2(folder []string) []string {
	var trie = make([][26]int, 100100)
	var index = 0
	var existValues = make(map[int]interface{})
	for i := 0; i < len(folder); i++ {
		var str = folder[i]
		var nextIndex = 0
		for j := 0; j < len(folder[i]); j++ {
			if str[j] == '/' {
				continue
			}
			var temp = str[j] - 'a'
			if trie[nextIndex][temp] == 0 {
				index++
				trie[nextIndex][temp] = index
			}
			nextIndex = index
		}
		existValues[nextIndex] = nil
	}
	var result = make([]string, 0, 10010)
	for i := 0; i < len(folder); i++ {
		var str = folder[i]

		var nextPointer = 0
		var ending = len(str)
		//lastPre
		var hasFind = false
		for pointer := len(str) - 1; pointer >= 0; pointer-- {
			if str[pointer] == '/' {
				if pointer == 0 {
					hasFind = true
				}
				ending = pointer
				break
			}
		}
		if hasFind {
			result = append(result, str)
			continue
		}
		for j := 0; j < ending; j++ {
			if str[j] == '/' {
				continue
			}
			var temp = str[j] - 'a'
			var trieValues = trie[nextPointer][temp]
			if _, ok := existValues[trieValues]; trieValues != 0 && ok {
				break
			}
			if j == ending-1 {
				result = append(result, str)
			}
			nextPointer = trie[nextPointer][temp]
		}
	}
	return result
}

func removeSubfolders(folder []string) []string {
	var result = make([]string, 0, 10010)
	var mapping = make(map[string]interface{})
	for i := 0; i < len(folder); i++ {
		mapping[folder[i]] = nil
	}
	for i := 0; i < len(folder); i++ {
		if !matchPre(folder[i], mapping) {
			result = append(result, folder[i])
		}
	}
	return result
}

func matchPre(name string, mapping map[string]interface{}) bool {
	var index = 1
	for index < len(name) {
		var findNextIndex = index
		for i := findNextIndex; i < len(name); i++ {
			if name[i] == '/' {
				findNextIndex = i
				break
			}
			if i == len(name)-1 {
				return false
			}
		}
		var temp = name[:findNextIndex]
		if _, ok := mapping[temp]; ok {
			return true
		}
		index = findNextIndex + 1
	}
	return false
}
