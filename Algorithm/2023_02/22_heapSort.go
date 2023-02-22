package main

import (
	"bufio"
	"fmt"
	"os"
)

var heap []int
var tail int

func main() {
	reader := bufio.NewReader(os.Stdin)
	var n, m int
	fmt.Fscanln(reader, &n, &m)
	heap = make([]int, n)
	tail = -1
	for i := 0; i < n; i++ {
		var element int
		fmt.Fscan(reader, &element)
		push(element)
	}
	for i := 0; i < m; i++ {
		fmt.Printf("%d ", pop())
	}
}
func swap(x int, y int) {
	if heap[x] == heap[y] {
		return
	}
	heap[x] = heap[x] ^ heap[y]
	heap[y] = heap[x] ^ heap[y]
	heap[x] = heap[x] ^ heap[y]
}
func down(location int) {
	var nextLocation int = location
	if location*2 <= tail && heap[nextLocation] > heap[location*2] {
		nextLocation = location * 2
	}
	if location*2+1 <= tail && heap[nextLocation] > heap[location*2+1] {
		nextLocation = location*2 + 1
	}
	if location != nextLocation {
		swap(location, nextLocation)
		down(nextLocation)
	}

}
func up(location int) {
	if location == 0 {
		return
	}
	if location/2 >= 0 && heap[location/2] > heap[location] {
		swap(location, location/2)
		up(location / 2)
	}
}
func push(x int) {
	tail++
	heap[tail] = x
	up(tail)
}
func pop() int {
	var result = heap[0]
	heap[0] = heap[tail]
	down(0)
	return result
}
