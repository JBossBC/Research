package main

//
//import (
//	"bufio"
//	"fmt"
//	"os"
//)
//
//var collect = make([]int, 100100)
//
//func find(nums int) int {
//	if collect[nums] != nums {
//		collect[nums] = find(collect[nums])
//	}
//	return collect[nums]
//}
//
//var result = make([]string, 0, 100100)
//
//func main() {
//	reader := bufio.NewReader(os.Stdin)
//	var operationNums int
//	var numsLen int
//	fmt.Fscanln(reader, &numsLen, &operationNums)
//	for i := 0; i < numsLen; i++ {
//		collect[i] = i
//	}
//	for i := 0; i < operationNums; i++ {
//		var operate byte
//		var A, B int
//		fmt.Fscanf(reader, "%c %d %d\n", &operate, &A, &B)
//		switch operate {
//		case 'M':
//			collect[find(A)] = find(B)
//		case 'Q':
//			if find(A) == find(B) {
//				result = append(result, "Yes")
//			} else {
//				result = append(result, "No")
//			}
//		}
//	}
//	for i := 0; i < len(result); i++ {
//		fmt.Println(result[i])
//	}
//}
