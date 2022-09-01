package main

func main() {

}

func shuffle(nums []int, n int) []int {
	var result = make([]int, len(nums))
	var pointer int
	var i, j int = 0, n
	for i < n {
		result[pointer] = nums[i]
		pointer++
		result[pointer] = nums[j]
		pointer++
		i++
		j++
	}
	return result
}
