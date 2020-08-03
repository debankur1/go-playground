package main

import (
	"container/list"
	"fmt"
	"sync"
	"time"
)
type PubSubData struct {
	lock sync.RWMutex
	List list.List
	MaxItem int
	MinSize int
	cond *sync.Cond
}
func (pubSub *PubSubData) Init(maxItem int){
	   pubSub.cond = sync.NewCond(&pubSub.lock)
	   pubSub.MaxItem = maxItem
}
func (pubSub *PubSubData) Publish(data int){
   pubSub.lock.Lock()
   for pubSub.List.Len() == pubSub.MaxItem{
   	  fmt.Println("List is full waiting....")
   	  pubSub.cond.Wait()
   }
   pubSub.List.PushFront(data)
   pubSub.cond.Broadcast()
   pubSub.lock.Unlock()
}
func (pubSub *PubSubData) Consume(){
	pubSub.lock.RLock()
	fmt.Println(pubSub.List.Front().Value)
	//pubSub.cond.Broadcast()
	pubSub.lock.RUnlock()
}
func main(){
	pubSub := PubSubData{}
	pubSub.Init(4)
	for i:=0;i<100;i++{
		go pubSub.Publish(20 + i)
		go pubSub.Consume()
		go pubSub.Consume()
		time.Sleep(2 *time.Second)
	}

}