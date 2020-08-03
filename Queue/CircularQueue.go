package main

import "fmt"

type CircularQueue struct {
	Q []int
	Front int
	Rear int
}
func (Q *CircularQueue) NewQueue(size int) *CircularQueue {
	var q = make([]int,size)
	var front int = 0
	var rear  int = 0
	return &CircularQueue{Q: q, Front: front, Rear: rear}
}
func(Q *CircularQueue) Enqueue(data int){
	if Q.Q!=nil{
		if !Q.isFull(){
			Q.Rear = (Q.Rear + 1)% len(Q.Q)
			Q.Q[Q.Rear]=data
		}
	} else{
		fmt.Println("Please intialize the queue calling the New Queue")
	}
}
func(Q *CircularQueue)Deqeue() int{
	if Q.Q!=nil {
		if !Q.IsEmpty() {
			Q.Front = (Q.Front + 1)% len(Q.Q)
			result := Q.Q[Q.Front]
			Q.Q[Q.Front] = 0
			return result
		}
	} else {
		fmt.Println("Please intialize the queue calling the New Queue")
	}
	return -1
}
func (Q *CircularQueue) IsEmpty() bool{
	if Q.Front >= Q.Rear{
		if Q.Q[Q.Front]==0{
			Q.Rear =  0
			Q.Front = 0
		}
		return true
	}
	return false
}
func (Q *CircularQueue) isFull() bool{
	if (Q.Rear + 1)% len(Q.Q)==Q.Front{
		return true
	}
	return false
}
