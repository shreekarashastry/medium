package main

import "fmt"

type Node struct{
	next  *Node
	value int
}

type LinkedList struct{
	head *Node
	len  int
}

func (L *LinkedList) Insert(value int){
	// create a new Node to insert the data
	newNode := Node{}

	newNode.value = value
	
	// If the list is empty
	if L.len == 0 {
		L.head = &newNode
		L.len++
		return			
	}
	
	// If the list is not empty
	currentHead := L.head
	
	for i:=0;i<L.len;i++ {
		if currentHead.next == nil {
			currentHead.next = &newNode
			L.len++
			return
		}
		currentHead = currentHead.next		
	} 
}

func (L *LinkedList) Display(){

	if L.len == 0 {
		fmt.Println("Empty List")
	}

	currentHead := L.head
	for i:=0;i<L.len;i++{
		fmt.Printf("%d ->", currentHead.value)
		currentHead = currentHead.next 	
	}
}

func main() {
	l := LinkedList{}	
	l.Display()

	l.Insert(10)
	l.Insert(5)
	l.Display()
}
