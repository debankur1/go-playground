package main

import (
	"fmt"
	"runtime"
	"time"
)
//Switching Context using runtime.Gosched
func Speak(s string){
	for i:=0;i<1000;i++{
		runtime.Gosched()
		fmt.Println(s)
	}
}
func main()  {
	 runtime.GOMAXPROCS(2)
	 go Speak("Hi")
	 Speak("Hello")
	 go Speak("Bye")
	time.Sleep(20 * time.Second)
}