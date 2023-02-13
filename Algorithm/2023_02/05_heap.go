//package main
//
//import (
//	"bufio"
//	"fmt"
//	"os"
//)
//
//var heap []int = make([]int, 100100)
//
////point the last element
//var heapLength int = 0
//
////第几个插入的元素 <-> 堆里面的下标
//var insertValue = make([]int, 10010)
//var pointer = 0
//
////堆里面的下标<->第几个插入的元素
//var reverseValue = make([]int, 100010)
//
//func main() {
//	var result = make([]int, 0, 30)
//	var operateNums int
//	reader := bufio.NewReader(os.Stdin)
//	fmt.Fscanln(reader, &operateNums)
//	var operate string
//	var index, number int
//	for i := 0; i < operateNums; i++ {
//		fmt.Fscan(reader, &operate)
//		switch operate {
//		case "I":
//			fmt.Fscanln(reader, &index)
//			heapLength++
//			insertValue[pointer] = heapLength
//			reverseValue[heapLength] = pointer
//			pointer++
//			heap[heapLength] = index
//			up(heapLength)
//		case "D":
//			fmt.Fscanln(reader, &index)
//			heapIndex := insertValue[index-1]
//			heap_swap(heapIndex, heapLength)
//			heapLength--
//			down(heapIndex)
//			up(heapIndex)
//		case "C":
//			fmt.Fscanln(reader, &index, &number)
//			heapValue := insertValue[index-1]
//			heap[heapValue] = number
//			down(heapValue)
//			up(heapValue)
//		case "PM":
//			fmt.Fscanln(reader)
//			result = append(result, heap[1])
//		case "DM":
//			fmt.Fscanln(reader)
//			heap_swap(1, heapLength)
//
//			heapLength--
//			down(1)
//		}
//	}
//	for i := 0; i < len(result); i++ {
//		fmt.Println(result[i])
//	}
//}
//func up(index int) {
//	for index/2 != 0 && heap[index] < heap[index/2] {
//		heap_swap(index, index/2)
//		index = index / 2
//	}
//}
//
//func down(index int) {
//	var nextSwap int = index
//	if index*2 <= heapLength && heap[nextSwap] > heap[2*index] {
//		nextSwap = 2 * index
//	}
//	if index*2+1 <= heapLength && heap[nextSwap] > heap[2*index+1] {
//		nextSwap = 2*index + 1
//	}
//	if nextSwap != index {
//		heap_swap(nextSwap, index)
//		down(nextSwap)
//	}
//}
//
//func heap_swap(pre int, last int) {
//	if pre == last {
//		return
//	}
//	insertValue[reverseValue[pre]] = insertValue[reverseValue[pre]] ^ insertValue[reverseValue[last]]
//	insertValue[reverseValue[last]] = insertValue[reverseValue[pre]] ^ insertValue[reverseValue[last]]
//	insertValue[reverseValue[pre]] = insertValue[reverseValue[pre]] ^ insertValue[reverseValue[last]]
//	reverseValue[pre] = reverseValue[pre] ^ reverseValue[last]
//	reverseValue[last] = reverseValue[pre] ^ reverseValue[last]
//	reverseValue[pre] = reverseValue[pre] ^ reverseValue[last]
//	heap[pre] = heap[pre] ^ heap[last]
//	heap[last] = heap[pre] ^ heap[last]
//	heap[pre] = heap[pre] ^ heap[last]
//}
