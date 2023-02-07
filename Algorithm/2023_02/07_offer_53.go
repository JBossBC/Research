package main

func main() {
}
func missingNumber(nums []int) int {
	var search = make(map[int]interface{})
	for i := 0; i < len(nums); i++ {
		search[nums[i]] = nil
	}
	for i := 0; i <= len(nums); i++ {
		if _, ok := search[i]; !ok {
			return i
		}
	}
	return -1
}
