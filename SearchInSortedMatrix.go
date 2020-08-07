package main

import "fmt"

func main(){
	mat := [][]int{
		{1,4,7,12,   15,1000},
		{2,5,19,31,  32,1001},
		{3,8,24,33,  35,1002},
		{40,41,42,44,45,1003},
		{99,100,103,106,128,1004},
	}
	targetNum := 103
	fmt.Println(SearchSortedMatrix(mat,targetNum))
}
//Time Complexity O(Row + Column) and Space Complexity is O(1)
func SearchSortedMatrix(matrix [][]int,targetNum int)[]int{
	row := 0
	col := len(matrix[0])-1
	for row<len(matrix) && col>=0{
		if matrix[row][col] >targetNum{
			col -= 1
		}else if matrix[row][col] < targetNum{
			row++
		}else{
			return []int{row,col}
		}
	}
	return []int{}
}