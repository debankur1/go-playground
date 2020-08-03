package main

import (
	"fmt"
	"go-jwt/Constructors"
)
func main(){
	Employee := Constructors.NewEmployee("a","2")
	fmt.Println(Employee)
}
