package main

import (
	"fmt"
	"sync"
	"time"
)
type MyMap struct {
	safe_map map[int]int
	sync.RWMutex
}
func runWriters(Map *MyMap,n int){
	for i:=0;i<20;i++{
		  Map.Lock()
		  Map.safe_map[i] = i + n
		  fmt.Println("From writer.....",i)
		  Map.Unlock()
		  time.Sleep(1 * time.Second)
	}
}
func runReaders(Map *MyMap,n int,whichReader string){
	for i:=0;i<10;i++{
		Map.RLock()
		n = n +1
		v := Map.safe_map[i]
		Map.RUnlock()
		time.Sleep(1 * time.Second)
		fmt.Println("From reader " + whichReader,v)
	}
}
func main(){
	myMap := MyMap{}
	myMap.safe_map = make(map[int]int)
	for {
		go runWriters(&myMap, 10)
		go runReaders(&myMap, 10, "1")
		go runReaders(&myMap, 10, "2")
		time.Sleep(20 * time.Second)
	}
}