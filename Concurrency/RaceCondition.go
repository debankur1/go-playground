package main

import "sync"

type test_counter struct {
	counterLock sync.Mutex
	count int
}
func (testCounter *test_counter) Increment(){
	testCounter.counterLock.Lock()
	testCounter.count ++
	testCounter.counterLock.Unlock()
}
func (testCounter *test_counter) Decrement(){
	testCounter.counterLock.Lock()
	testCounter.count --
	testCounter.counterLock.Unlock()
}
func (testCounter *test_counter)GetValue()int{
	testCounter.counterLock.Lock()
	v:= testCounter.count
	testCounter.counterLock.Unlock()
	return v
}
func main(){
	testCounter := new(test_counter)
	for i:= 0;i<10;i++{
		 go testCounter.Increment()
		 go testCounter.Decrement()
	}
	testCounter.GetValue()
}