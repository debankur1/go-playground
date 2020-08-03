package main

import (
	"fmt"
	"sync"
	"time"
)
var mutex sync.Mutex
var wait sync.WaitGroup
func main(){
	//oddChan := make(chan bool)
	//evenChan := make(chan bool)
    //wait.Add(2)
	//go PrintOdd(oddChan,evenChan)
	//go PrintEven(oddChan,evenChan)
    //wait.Wait()

	//wait.Add(2)
	//p := make(chan int)
	//go PrintEventChan(p)
	//go PrintOddOneChan(p)
	//wait.Wait()
	cond := sync.NewCond(&sync.Mutex{})
	go func(){
		for i:=0;i<10;i++{
			if i%2==0{
				cond.L.Lock()
				fmt.Println("Even :",i)
				cond.L.Unlock()
			} else{
				cond.Broadcast()
			}
		}

	}()
	go func(){
		for i:=0;i<10;i++{
			if i%2!=0{
				cond.L.Lock()
				fmt.Println("Odd :",i)
				cond.L.Unlock()
			}else{
				cond.Broadcast()
			}
		}

	}()

time.Sleep(3 * time.Second)
}
func PrintOdd(odd chan bool,even chan bool){
 for i :=0;i<10;i++{
 	if i%2!=0{
		<- even
 		fmt.Println("odd Number",i)
 		odd <- true
	}
 }
 defer wait.Done()
}
func PrintEven(odd chan bool,even chan bool) {
	for i :=0;i<10;i++{
		if i%2==0{
			fmt.Println("Even Number",i)
			even <- true
			<- odd
		}
	}
	defer wait.Done()
}

func PrintOddOneChan(p chan int){
	for i:= 1;i<10;i++{
		p <-i
		if i%2!=0{
			fmt.Println("Odd Number",i)
		}
	}
	defer wait.Done()
}

func PrintEventChan(p chan int){
	for i:= 1;i<10;i++{
		<-p
		if i%2 == 0{
			fmt.Println("Even Number",i)
		}
	}
	defer wait.Done()
}
