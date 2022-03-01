package main

import (
	"fmt"
	"goSinglyLinkedList/singlylinkedlist"
)

func main() {
	list := singlylinkedlist.InitList()

	// insert values at front
	list.InsertFront(5)
	list.InsertFront(6)
	list.InsertFront(7)
	list.InsertFront(8)

	// insert values at back
	list.InsertBack(10)
	list.InsertBack(20)
	list.InsertBack(30)

	// insert values at specific position
	err := list.InsertAt(100, 3)
	if err != nil {
		fmt.Println("Error: ", err.Error())
	}

	list.Display()
	//8 7 6 100 5 10 20 30
	err = list.RemoveFront()
	if err != nil {
		fmt.Println("Error: ", err.Error())
	}

	list.Display()
	//7 6 100 5 10 20 30

	err = list.RemoveBack()
	if err != nil {
		fmt.Println("Error: ", err.Error())
	}

	list.Display()
	//7 6 100 5 10 20

	err = list.RemoveAt(3)
	if err != nil {
		fmt.Println("Error: ", err.Error())
	}

	list.Display()
	//7 6 100 10 20
}
