package main
import (
	"sort"
	"strings"
	"fmt"
)
func main(){
   var result=stringMatching([]string{"mass","as","hero","superhero"})
   for _,value:=range result{
	fmt.Printf("%9s",value)
   }
}
func stringMatching(words []string) []string {
	if len(words)<=1{
		return nil
	}
	sort.Slice(words,func(one,two int)bool{
                 return len(words[two])<len(words[one])
	})
	var tempPointer=-1
	var result []string
	for i:=0;i<len(words);i++{
		tempPointer=i-1
		for tempPointer>=0 {
			if strings.Contains(words[tempPointer],words[i]){
                result=append(result,words[i])
				break
			}
			tempPointer--
		}
	}
	return result
}