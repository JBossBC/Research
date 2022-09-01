package main

func main() {

}

type OrderedStream struct {
	arr     []string
	pointer int
}

func Constructor(n int) OrderedStream {
	return OrderedStream{
		arr:     make([]string, n),
		pointer: 0,
	}
}

func (this *OrderedStream) Insert(idKey int, value string) []string {
	this.arr[idKey-1] = value
	var index = idKey - 1
	if this.pointer >= len(this.arr) {
		return nil
	}
	if this.pointer != index {
		return nil
	}
	for index < len(this.arr) && this.arr[index] != "" {
		index++
	}
	this.pointer = index
	return this.arr[idKey-1 : index]
}

/**
 * Your OrderedStream object will be instantiated and called as such:
 * obj := Constructor(n);
 * param_1 := obj.Insert(idKey,value);
 */
