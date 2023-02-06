package main

func main() {

}
func findRepeatNumber(nums []int) int {
	var repeatMapping = make(map[int]int)
	for i := 0; i < len(nums); i++ {
		repeatMapping[nums[i]]++
		if repeatMapping[nums[i]] > 1 {
			return nums[i]
		}
	}
	return -1
}
