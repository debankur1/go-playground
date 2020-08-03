package main

import "fmt"

// Time Complexity is O(N) | Space O(1)
func ReverseString(s string)string{
	strArr := make([]byte, 0, len(s)) // Slices has property to add Length and Capacity as total size
	for k:=0;k< len(s);k++{
		strArr = append(strArr, s[k])
	}
	i := 0
	j:= len(strArr)
	for i<j{
		j = j-1
		x := strArr[i]
		strArr[i] = strArr[j]
		strArr[j] = x
		i=i+1
	}
	return string(strArr)
}
func CheckPalindrome(s1 string) bool{
	if s1 == ReverseString(s1){
		return true
	}
	return false
}
func main(){
	fmt.Println(CheckPalindrome("madam"))
}