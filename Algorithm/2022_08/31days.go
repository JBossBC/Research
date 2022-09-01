package main

func main() {

}
func validateStackSequences(pushed []int, popped []int) bool {
	var stack = make([]int, len(pushed))
	var stackPointer = -1
	var pushedIndex = 0
	var poppedIndex = 0
	for pushedIndex < len(pushed) {
		stackPointer++
		stack[stackPointer] = pushed[pushedIndex]
		for stackPointer >= 0 && popped[poppedIndex] == stack[stackPointer] {
			stackPointer--
			poppedIndex++

		}
		pushedIndex++
	}
	if poppedIndex >= len(popped) {
		return true
	}
	return false
}
