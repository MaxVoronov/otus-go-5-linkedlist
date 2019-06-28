package doublylinkedlist

import "fmt"

type List struct {
	head *Node
	tail *Node
}

type Node struct {
	Value interface{}
	Prev  *Node
	Next  *Node
}

func NewList() *List {
	return &List{}
}

func (lst *List) Len() int {
	counter := 0
	cursor := lst.head
	for cursor != nil {
		counter++
		cursor = cursor.Next
	}

	return counter
}

func (lst *List) First() *Node {
	return lst.head
}

func (lst *List) Last() *Node {
	return lst.tail
}

func (lst *List) PushFront(data interface{}) *Node {
	var node *Node
	if lst.head == nil {
		node = &Node{data, nil, nil}
		lst.tail = node
	} else {
		node = lst.PushBefore(lst.head, data)
	}
	lst.head = node

	return node
}

func (lst *List) PushBack(data interface{}) *Node {
	var node *Node
	if lst.tail == nil {
		node = &Node{data, nil, nil}
		lst.head = node
	} else {
		node = lst.PushAfter(lst.tail, data)
	}
	lst.tail = node

	return node
}

func (lst *List) PushBefore(node *Node, data interface{}) *Node {
	newNode := &Node{data, node.Prev, node}
	if node.Prev != nil {
		node.Prev.Next = newNode
	} else {
		lst.head = newNode
	}
	node.Prev = newNode

	return newNode
}

func (lst *List) PushAfter(node *Node, data interface{}) *Node {
	newNode := &Node{data, node, node.Next}
	if node.Next != nil {
		node.Next.Prev = newNode
	} else {
		lst.tail = newNode
	}
	node.Next = newNode

	return newNode
}

func (lst *List) Remove(node *Node) {
	if node.Prev != nil {
		node.Prev.Next = node.Next
	} else {
		lst.head = node.Next
	}

	if node.Next != nil {
		node.Next.Prev = node.Prev
	} else {
		lst.tail = node.Prev
	}
}

func (lst *List) String() string {
	result := fmt.Sprintf("[H] -> ")
	cursor := lst.head
	for cursor != nil {
		result += fmt.Sprintf("[%v] -> ", cursor.Value)
		cursor = cursor.Next
	}
	result += fmt.Sprintf("[T]")

	return result
}
