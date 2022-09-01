package main

func main(){
   print(busyStudent([]int{1,2,3},[]int{3,5,7},3))

}
func busyStudent(startTime []int, endTime []int, queryTime int) int {
	if len(startTime)!=len(endTime){
		return -1
	}
	var sum=0
	for i:=0;i<len(startTime);i++{
		 if startTime[i]<=queryTime&&endTime[i]>=queryTime{
			 sum++
		 }
	}
	return sum
}