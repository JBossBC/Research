package main

import(
	"fmt"
	"os"
	"bufio"
)
var hash []int
var list []int
var next int
var hashValue=10003
func main(){
	hash =make([]int,hashValue)
	list=make([]int,100010)
	next= 0
	reader:=bufio.NewReader(os.Stdin)
	var n int
	fmt.Fscanln(reader,&n)
    
}

func insert(x int){
    point:=hash[x%hashValue]
	if point == 0 
}