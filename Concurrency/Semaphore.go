package main

import (
	"context"
	"fmt"
	"golang.org/x/sync/semaphore"
	"time"
)
func main(){
	sem := semaphore.NewWeighted(4)
	ctx := context.Background()
	for i:=0;i<100;i++{
		 sem.Acquire(ctx,1)
		 go func(){
		 	fmt.Println("Executed by Semaphore")
		 	time.Sleep(2 * time.Second)
			 sem.Release(1)
		 }()
	}
	 CustomSemaphore(4)
}
func CustomSemaphore(weight int){
	sch := make(chan int,weight)
	for i:=0;i<100;i++{
		sch <- i
		go func(){
			fmt.Println("Execute form Custom Semaphore")
			time.Sleep(2 * time.Second)
			<- sch
		}()
	}
}