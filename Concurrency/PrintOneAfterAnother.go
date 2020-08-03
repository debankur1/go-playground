package main

import (
	"fmt"
)

var ch= make(chan bool)
var ch2 = make(chan bool)
func main() {
	for i:=0;i<100;i++ {
		wait.Add(2)
		go player1()
		go player2()
		//go player3()
		wait.Wait()
	}
}
func player1() {
	fmt.Println("Player 1")
	ch <- true
	wait.Done()
}
func player2(){
	if <-ch {
		fmt.Println("Player 2")
		//ch2 <- true
		wait.Done()
	}
}
func player3(){
	if <-ch2 {
		fmt.Println("Player 3")
		wait.Done()
	}

}

