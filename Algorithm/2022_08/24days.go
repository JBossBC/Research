package main

func main() {
	print(canBeEqual([]int{4, 3, 2, 1}, []int{1, 2, 3, 4}))
}

func canBeEqual(target []int, arr []int) bool {\

	var ValueMap = make(map[int]int, len(target))
	for i := 0; i < len(target); i++ {
		ValueMap[target[i]]++
	}
	for i := 0; i < len(arr); i++ {
		if _, ok := ValueMap[arr[i]]; ok {
			ValueMap[arr[i]]--
		} else {
			return false
		}
	}
	for _, value := range ValueMap {
		if value != 0 {
			return false
		}
	}
	return true

}
