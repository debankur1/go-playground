package main

import (
	"container/list"
	"fmt"
	"sync"
	"time"
)

var Lock sync.RWMutex
var maxsize int = 10
var minSize int = 0
var List list.List
var Cond  = sync.NewCond(&Lock)
var wg sync.WaitGroup
var value int = 0
func Publish(){
	defer wg.Done()
 for i:=0;i<200;i++{
 	Lock.Lock()
 	for List.Len()==10{
 		fmt.Println("List is full...Waiting")
 		Cond.Wait()
	}
	 value++
 	 fmt.Println("Prodcuing,",value)
 	 List.PushFront(value)
	 Cond.Broadcast()
 	 Lock.Unlock()
 }
}
func Consume(){
	defer wg.Done()
	for i:=0;i<100;i++ {
		Lock.RLock()
		for List.Len() == 0 {
			fmt.Println("List is empty...Waiting")
			Cond.Wait()
		}
		fmt.Println("Consuming",List.Front().Value)
		List.Remove(List.Front())
		Cond.Broadcast()
		Lock.RUnlock()
	}
}
func main(){
   // wg.Add(3)
	go Publish()
	go Consume()
    go Consume()
  //  wg.Wait()

  time.Sleep(20 * time.Second)

}