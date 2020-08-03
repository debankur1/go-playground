package main

import "fmt"

//Time Complexity O(n) | Space O(N=Number of Alphabets --26)
func FindAnagram(s1 string,s2 string)bool{
	if len(s1)!= len(s2){
		return false
	}
	HashTable := make([]int,26)
	for i:=0;i< len(s1);i++{
		index := s1[i]-97
		HashTable[index] += 1
	}
	for i:=0;i< len(s2);i++{
		index := s2[i]-97
		HashTable[index] -= 1
		if HashTable[index]==-1{
			return false
		}
	}
	return true
}
func main()  {
	fmt.Println(FindAnagram("decimal","medical"))


}