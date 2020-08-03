package main

import "fmt"

func main(){
	x := "Hello!!!!&&&&&%%%%__+++"
	runes := []rune(x)
	for i:=0;i<len(runes);i++{
		 fmt.Println(string(runes[i]))
	}
}
