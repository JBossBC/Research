package main

import (
	"fmt"
)

var e = make([]int, 10010)
var ne = make([]int, 10010)
var nextLocation int = 1

func Add(nums int) {
	ne[nextLocation] = ne[0]
	e[nextLocation] = nums
	ne[0] = nextLocation
	nextLocation++
}
func Delete(k int) {
	ne[k] = ne[ne[k]]
}
func Insert(k int, nums int) {
	e[nextLocation] = nums
	ne[nextLocation] = ne[k]
	ne[k] = nextLocation
	nextLocation++
}
func main() {
	var times int
	fmt.Scanln(&times)
	for i := 0; i < times; i++ {
		var operationName byte
		fmt.Scanf("%c", &operationName)
		switch operationName {
		case 'H':
			var values int
			fmt.Scanln(&values)
			Add(values)
		case 'I':
			var insertLocation int
			var values int
			fmt.Scanln(&insertLocation, &values)
			Insert(insertLocation, values)
		case 'D':
			var deleteLocation int
			fmt.Scanln(&deleteLocation)
			Delete(deleteLocation)
		}
	}

	for i := ne[0]; i != 0; i = ne[i] {
		fmt.Printf("%d ", e[i])
	}
}
