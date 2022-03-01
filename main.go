package main

import "fmt"

type Node struct {
	key  interface{}
	next *Node
}

type List struct {
	head *Node
	len  int
}

//initList returns an empty list
func initList() *List {
	return &List{}
}

//InsertFront inserts an element at the head of the list
func (l *List) InsertFront(key interface{}) {
	node := &Node{
		key:  key,
		next: nil,
	}

	if l.head == nil {
		l.head = node
	} else {
		node.next = l.head
		l.head = node
	}
	l.len++
	return
}

//InsertBack inserts an element at the tail of the list
func (l *List) InsertBack(key interface{}) {
	node := &Node{
		key:  key,
		next: nil,
	}

	if l.head == nil {
		l.head = node
	} else {
		current := l.head

		for current.next != nil {
			current = current.next
		}
		current.next = node
	}
	l.len++
	return
}

func (l *List) Display() {
	current := l.head

	for current != nil {
		fmt.Println("Node: ", current.key)
		current = current.next
	}

}

func main() {
	list := initList()

	list.InsertFront(5)
	list.InsertFront(6)
	list.InsertFront(7)
	list.InsertFront(8)

	//list.Display()

	list.InsertBack(10)
	list.InsertBack(20)
	list.InsertBack(30)

	list.Display()
}
