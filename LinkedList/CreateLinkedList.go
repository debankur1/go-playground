package main

import "fmt"

type Node struct {
	data int
	Next *Node
}
//O(1) Time Complexity
var First *Node
func InsertLinkedList(First *Node,Last *Node,data int) (*Node,*Node){
	newNode :=  &Node{
		data: data,
		Next: nil,
	}
	if First==nil{
		First = newNode
		Last  = newNode
	}else{
		Last.Next = newNode
		Last = newNode
	}
	return First,Last
}
func InsertAtaPositionInaLinkedList(head *Node,data int,position int) *Node{
	newNode := &Node{
		data: data,
		Next: nil,
	}
	temp := head
	var prev *Node
	for i:=0;i<=position;i++{
		prev = temp
		temp = temp.Next
	}
	prev.Next = newNode
	newNode.Next = temp

	return head
}
func ReverseLinkedList(LinkedList *Node)*Node{
	p := LinkedList
	var q *Node
	var r *Node
	for p!=nil{
		r = q
		q = p
		p = p.Next
		q.Next = r
	}
	return q
}
func FindMiddileNodeOfaLinkedList(LinkedList *Node) *Node{
	p,q := LinkedList,LinkedList
	for q!=nil{
		q = q.Next
		if q!=nil{
			q = q.Next
		}
		if q!=nil{
			p = p.Next
		}
	}
	return p
}
func MergeTwoLinkedList(l1 *Node,l2 *Node)*Node{
	third :=  &Node{}
	last  :=  &Node{}
	if l1.data < l2.data{
		third = l1
		last = l1
		l1 = l1.Next
		last.Next = nil
	} else {
		third = l2
		last = l2
		l2 = l2.Next
		last.Next = nil
	}
	for l1!=nil && l2!=nil{
		if l1.data < l2.data {
			last.Next = l1
			last= l1
			l1 = l1.Next
			last.Next = nil
		}else{
			last.Next = l2
			last= l2
			l2 = l2.Next
			last.Next = nil
		}
	}
	if l1!=nil{
		last.Next = l1
	} else {
		last.Next = l2
	}
return third
}
func main()  {
	var LinkedList *Node
	var Last *Node
	result,last :=  InsertLinkedList(LinkedList,Last,20)
	result,last  =  InsertLinkedList(result,last,25)
	result,last  =  InsertLinkedList(result,last,30)
	result,last  =  InsertLinkedList(result,last,40)
	result,last  =  InsertLinkedList(result,last,50)

	var LinkedList1 *Node
	var Last1 *Node
	result1,last1 :=  InsertLinkedList(LinkedList1,Last1,60)
	result1,last1  =  InsertLinkedList(result1,last1,70)
	result1,last1  =  InsertLinkedList(result1,last1,80)
	result1,last1  =  InsertLinkedList(result1,last1,90)
	result1,last1  =  InsertLinkedList(result1,last1,100)
    //result = FindMiddileNodeOfaLinkedList(result)

    mergeResult := MergeTwoLinkedList(result,result1)
    fmt.Println(mergeResult)

	//result = ReverseLinkedList(result)
	//First = result
	//result = InsertAtaPositionInaLinkedList(result,34,1)
}