package main

import "fmt"

type Queue struct {
	Q []int
	Front int
	Rear int
}
func (Q *Queue) NewQueue(size int) *Queue {
	var q = make([]int,size)
	var front int = -1
	var rear  int = -1
	return &Queue{Q: q, Front: front, Rear: rear}
}
func(Q *Queue) Enqueue(data int){
	if Q.Q!=nil{
		if !Q.isFull(){
			Q.Rear = Q.Rear + 1
			Q.Q[Q.Rear]=data
		}
	} else{
		fmt.Println("Please intialize the queue calling the New Queue")
	}
}
func(Q *Queue)Deqeue(){
	if Q.Q!=nil {
		if !Q.IsEmpty() {
			Q.Front = Q.Front + 1
			Q.Q[Q.Front] = 0
		}
	} else {
		fmt.Println("Please intialize the queue calling the New Queue")
	}
}
func (Q *Queue) IsEmpty() bool{
	if Q.Front >= Q.Rear{
		if Q.Q[Q.Front]==0{
			 Q.Rear =  -1
			 Q.Front = -1
		}
		return true
	}
	return false
}
func (Q *Queue) isFull() bool{
	if Q.Rear== len(Q.Q)-1{
		return true
	}
	return false
}
func main() {
	Q := Queue{}
	queue :=Q.NewQueue(30)
	queue.Enqueue(10)
	queue.Enqueue(20)
	queue.Enqueue(30)
	queue.Enqueue(40)
	queue.Deqeue()
	queue.Deqeue()
	queue.Deqeue()
	queue.Deqeue()
}