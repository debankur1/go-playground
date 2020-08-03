package main

import (
	"fmt"
	"time"
)

func main(){
	c1 := make(chan int)
	c2 := make(chan int)

	for i:=1;i<10;i++{
		go func() {
			time.Sleep(1 * time.Second)
			c1 <- i
			c2 <- i+2
		}()
	}


	for i := 0; i < 10; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println("received", msg1)
		case msg2 := <-c2:
			fmt.Println("received", msg2)
		}
	}
}
