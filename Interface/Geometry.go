package main

import (
	"fmt"
	"reflect"
)

type Rectangle struct {
	w,l int
}
type Shape interface {
	area() int
}
func(r *Rectangle) area() int{
	return r.l * r.w
}
func(s *Square) area() int{
	return s.l * s.l
}
type Square struct {
	l int //Sides
}
func Info(s Shape){
	fmt.Println("the area is",s.area())
}
func main(){
	var Shapes [2]Shape
	Shapes[0] = &Rectangle{
		w: 10,
		l: 10,
	}
	Shapes[1] = &Square{
		l: 20,
	}

    Info(Shapes[0])
	Info(Shapes[1])

	//Interface assertion
	var message interface{} = 1924343
	s,ok := message.(string)
	fmt.Println(s,ok)

	x :=reflect.ValueOf(message)
	fmt.Println(x.Type())
}