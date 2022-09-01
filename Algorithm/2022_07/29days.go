package main

func main() {
}

func validSquare(p1 []int, p2 []int, p3 []int, p4 []int) bool {
	var isSquare = true
	var arr = [4][]int{p1, p2, p3, p4}
	var resumeMiddle = make([]int, 0)
	if p1[0] == p2[0] || p1[1] == p2[1] {
		if p1[1] == p3[1] {
			resumeMiddle = p1
		} else {
			resumeMiddle = p2
		}
	} else {
		resumeMiddle = p3
	}
	return false
}
