package main

import (
	"bufio"
	"fmt"
	"os"
)

var arr []int

//
//func down(index int) {
//	var swapLocation int = index
//	if index*2 < len(arr) && arr[swapLocation] >= arr[index*2] {
//		swapLocation = index * 2
//	}
//	if index*2+1 < len(arr) && arr[swapLocation] >= arr[index*2+1] {
//		swapLocation = index*2 + 1
//	}
//	//  if not stable
//	if swapLocation != index {
//		arr[index] = arr[index] ^ arr[swapLocation]
//		arr[swapLocation] = arr[index] ^ arr[swapLocation]
//		arr[index] = arr[index] ^ arr[swapLocation]
//		down(swapLocation)
//	}
//}

func main() {
	var arrLength, inputLength int
	reader := bufio.NewReader(os.Stdin)
	fmt.Fscanf(reader, "%d %d\n", &arrLength, &inputLength)
	arr = make([]int, arrLength+1)
	for i := 1; i < len(arr); i++ {
		fmt.Fscan(reader, &arr[i])
	}
	for i := (len(arr) - 1) / 2; i > 0; i-- {
		down(i)
	}
	var vaildLength = len(arr) - 1
	for i := 0; i < inputLength; i++ {
		fmt.Printf("%d ", arr[1])
		arr[1] = arr[vaildLength]
		vaildLength--
		down(1)
	}

}
