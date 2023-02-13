package main

import "fmt"

func main() {
	g := initMap()
	g.Insert(10005, 1)
	g.Insert(10004, 2)
	g.Insert(10003, 3)
	g.Insert(10002, 4)
	g.Insert(10001, 5)
	g.Insert(10000, 6)
	g.Insert(10000, 7)
	g.Insert(10000, 8)
	value, ok := g.Find(10000)
	if ok {
		fmt.Println(value)
	}
	value, ok = g.Find(10001)
	if ok {
		fmt.Println(value)
	}

}
func initMap() golangMap {
	return golangMap{
		header: make([]int, 100010),
		n:      make([]int, 100010),
		valuen: make([]int, 100010),
		ne:     make([]int, 100010),
		
		index: 1,
	}
}
func (m *golangMap) Insert(x int, value1 int) {
	yushu := x % 100
	if m.header[yushu] == 0 {
		m.n[m.index] = x
		m.valuen[m.index] = value1
		m.header[yushu] = m.index
		m.index++
	} else {
		var value = m.header[yushu]
		for m.ne[value] != 0 {
			if x == m.n[value] {
				m.valuen[value] = value1
				return
			}
			value = m.ne[value]
		}
		m.valuen[m.index] = value
		m.n[m.index] = x
		m.ne[value] = m.index
		m.index++
		//temp := value
		//m.valuen[m.index] = value
		//m.n[m.index] = x
		//m.ne[m.index] = temp
		//m.header[yushu] = m.index
		//m.index++
	}
}
func (m *golangMap) Find(x int) (int, bool) {
	yushu := x % 100
	value := m.header[yushu]
	for value != 0 {
		if m.n[value] == x {
			return m.valuen[value], true
		}
		value = m.ne[value]
	}
	return 0, false
}
func (golangMap) Delete(x int) {

}

type golangMap struct {
	header []int
	n      []int
	ne     []int
	valuen []int
	index  int
}
