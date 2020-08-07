package main

import "fmt"

func main(){
	x := []int{1,2,3,4}
    fmt.Println(GetPowerSet(x))
}
func GetPowerSet(set []int)[][]int{
	subsets := [][]int{{}}
	for _,ele := range set{
		length := len(subsets)
		for j:=0;j<length;j++{
			 currentSubset := subsets[j]
			 newSubset := append([]int{},currentSubset...)
			 newSubset  = append(newSubset,ele)
			 subsets    = append(subsets,newSubset)
		}
	}
	return subsets
}