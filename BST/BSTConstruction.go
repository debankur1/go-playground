package main

import "fmt"

type BST struct {
	data   int
	left   *BST
	Right  *BST
}
func (tree *BST) Insert(value int) *BST{
	current := tree
	for{
		if value < current.data{
			if current.left==nil{
				current.left = &BST{
					data: value,
				}
				break
			} else {
				current = current.left
			}
		} else {
			if current.Right==nil{
				current.Right = &BST{
					data: value,
				}
				break
			} else {
				current = current.Right
			}
		}
	}
	return tree
}
func (tree *BST) Contains(value int)bool{
	current := tree
	for current!=nil{
		 if value > current.data{
		 	   current = current.Right
		 } else if value<current.data{
		 	current = current.left
		 }else{
		 	return true
		 }
	}
   return false
}
func main()  {
	bst := BST{}
	bst.Insert(10)
	bst.Insert(20)
	bst.Insert(25)
	bst.Insert(30)
	bst.Insert(40)
	bst.Insert(35)
	bst.Insert(5)
	bst.Insert(8)
	bst.Insert(6)
	bst.Insert(9)

	fmt.Println(bst.Contains(80))

	var x string = "123"
	var y interface{}
	y = x
	fmt.Print(y)

}