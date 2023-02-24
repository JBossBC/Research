package main

func main() {

}

func hammingWeight(num uint32) int {
	var result int = 0
	for num != 0 {
		if num&1 != 0 {
			result++
		}
		num = num >> 1
	}
	return result
}
