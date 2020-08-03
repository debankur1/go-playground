package main

import (
	"container/list"
	"fmt"
)
type Node struct {
	data int
	key  string
}
type LRU struct {
	counter   int
	size       int
	doublyList *list.List
	CacheMap map[string]*list.Element
}
func (lru *LRU) Initialize(size int){
	lru.size = size
	lru.doublyList = new(list.List)
	lru.CacheMap = make(map[string]*list.Element)
}
func (lru *LRU) Add(node Node){
	if lru.counter <= lru.size {
		lru.doublyList.PushFront(node)
		dataInList := &list.Element{
			Value: node,
		}
		lru.CacheMap[node.key] = dataInList
		lru.counter++
	}else{
		//Accessing the node at the tail pointer of the LinkedList
		var delNode Node
		delNode = lru.doublyList.Back().Value.(Node)
		lru.doublyList.Remove(lru.doublyList.Back())
		delete(lru.CacheMap, delNode.key)
		lru.counter--
		lru.Add(node)
	}
}
func (lru *LRU)Get(key string) int{
	 requestedNode := lru.CacheMap[key]
	 lru.doublyList.Remove(requestedNode)
	 lru.doublyList.PushFront(requestedNode.Value)
	 val :=requestedNode.Value.(Node)
	 delete(lru.CacheMap, key)
	 lru.CacheMap[key] = requestedNode
	 return val.data
}
func main(){
	cache :=LRU{}
	cache.Initialize(3)
	cache.Add(Node{
		data: 10,
		key:  "a",
	})
	cache.Add(Node{
		data: 20,
		key:  "b",
	})
	cache.Add(Node{
		data: 20,
		key:  "c",
	})
	cache.Add(Node{
		data: 40,
		key:  "d",
	})
	cache.Add(Node{
		data: 50,
		key:  "e",
	})
   fmt.Println(cache.Get("c"))
	x := cache.doublyList.Front()
	fmt.Println(x)
}