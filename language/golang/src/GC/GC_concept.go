package main

import (
	"fmt"
)

package main

import (
"fmt"
)

func main() {
	fmt.Println("start.")

	container := make([]int, 8)
	fmt.Println("> loop.")
	for i := 0; i < 32*1000*1000; i++ {
		container = append(container, i)
	}
	fmt.Println("< loop.")
}
