package main

import (
	"fmt"
	"sync"
)
var wait sync.WaitGroup
func main(){
	oddChan,evenChan := make(chan bool),make(chan bool)
	wait.Add(2)
		go PrintHello(oddChan,evenChan)
	    go PrintHI(oddChan,evenChan)
	wait.Wait()
}
func PrintHI(odd chan bool,even chan bool){
	for {
			<- even
			fmt.Println("Hello")
			odd <- true
	}
	defer wait.Done()
}
func PrintHello(odd chan bool,even chan bool) {
	for {
			fmt.Println("Hi")
			even <- true
			<- odd
		}

	defer wait.Done()
}