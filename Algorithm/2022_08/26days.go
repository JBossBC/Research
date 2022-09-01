package main

func main() {

}
func maxProduct(nums []int) int {
	var maxIndex, secondMaxIndex int = 0, 0
	for index, Value := range nums {
		if Value >= nums[maxIndex] {
			maxIndex = index
		}
	}
	if maxIndex == 0 {
		secondMaxIndex++
	}
	for index, Value := range nums {
		if index == maxIndex {
			continue
		}
		if Value > nums[secondMaxIndex] {
			secondMaxIndex = index
		}
	}

	return (nums[maxIndex] - 1) * (nums[secondMaxIndex] - 1)
}
