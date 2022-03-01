package singlylinkedlist

import "fmt"

type Node struct {
	key  interface{}
	next *Node
}

type List struct {
	head *Node
	len  int
}

//InitList returns an empty list
func InitList() *List {
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

//InsertAt inserts an element at a specific position of the list (Zero based indexing)
func (l *List) InsertAt(key interface{}, pos int) error {
	if pos < 0 || pos > (l.len-1) {
		return fmt.Errorf("Out of bounds")
	}

	node := &Node{
		key:  key,
		next: nil,
	}

	// insert at front
	if pos == 0 {
		node.next = l.head
		l.head = node
	} else {
		current := l.head
		i := 0

		//find previous element
		for i < pos-1 {
			i++
			current = current.next
		}

		// add after
		node.next = current.next
		current.next = node
	}
	l.len++
	return nil
}

//RemoveFront removes head of list
func (l *List) RemoveFront() error {
	if l.head == nil {
		return fmt.Errorf("Empty list.")
	} else {
		current := l.head.next
		l.head = current
	}
	l.len--
	return nil
}

//RemoveBack removes tail of list
func (l *List) RemoveBack() error {
	if l.head == nil {
		return fmt.Errorf("Empty list.")
	} else {
		var previous *Node
		current := l.head
		for current.next != nil {
			previous = current
			current = current.next
		}

		if previous == nil { // one element list
			l.head = nil
		} else {
			previous.next = nil
		}
	}
	l.len--
	return nil
}

//RemoveAt removes an element at a specific position of the list (Zero based indexing)
func (l *List) RemoveAt(pos int) error {
	if pos < 0 || pos > (l.len-1) {
		return fmt.Errorf("Out of bounds")
	}

	current := l.head
	i := 0

	if pos == 0 {
		l.head = current.next
		l.len--
		return nil
	}

	// find previous element
	for i < pos-1 {
		i++
		current = current.next
	}

	removeNode := current.next
	current.next = removeNode.next
	l.len--
	return nil
}

func (l *List) Display() {
	current := l.head

	for current != nil {
		fmt.Print(current.key)
		fmt.Print(" ")
		current = current.next
	}
	fmt.Println()
}
