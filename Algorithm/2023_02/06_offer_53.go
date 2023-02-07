package main

func main() {

}
func search(nums []int, target int) int {
	var result int
	for i := 0; i < len(nums); i++ {
		if nums[i] == target {
			result++
		}
	}
	return result
}
