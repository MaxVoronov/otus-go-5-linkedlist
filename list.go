package doublylinkedlist

import "fmt"

type List struct {
	head   *Node
	tail   *Node
	length int
}

type Node struct {
	Value interface{}
	prev  *Node
	next  *Node
	list  *List
}

func NewList() *List {
	return &List{}
}

func (lst *List) Len() int {
	return lst.length
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
		node = &Node{data, nil, nil, lst}
		lst.tail = node
		lst.length++
	} else {
		node = lst.PushBefore(lst.head, data)
	}
	lst.head = node

	return node
}

func (lst *List) PushBack(data interface{}) *Node {
	var node *Node
	if lst.tail == nil {
		node = &Node{data, nil, nil, lst}
		lst.head = node
		lst.length++
	} else {
		node = lst.PushAfter(lst.tail, data)
	}
	lst.tail = node

	return node
}

func (lst *List) PushBefore(node *Node, data interface{}) *Node {
	newNode := &Node{data, node.prev, node, lst}
	if node.prev != nil {
		node.prev.next = newNode
	} else {
		lst.head = newNode
	}
	node.prev = newNode
	lst.length++

	return newNode
}

func (lst *List) PushAfter(node *Node, data interface{}) *Node {
	newNode := &Node{data, node, node.next, lst}
	if node.next != nil {
		node.next.prev = newNode
	} else {
		lst.tail = newNode
	}
	node.next = newNode
	lst.length++

	return newNode
}

func (lst *List) Remove(node *Node) {
	if node.list != lst {
		return
	}

	if node.prev != nil {
		node.prev.next = node.next
	} else {
		lst.head = node.next
	}

	if node.next != nil {
		node.next.prev = node.prev
	} else {
		lst.tail = node.prev
	}

	lst.length--
}

func (lst *List) String() string {
	result := fmt.Sprintf("[H] -> ")
	cursor := lst.head
	for cursor != nil {
		result += fmt.Sprintf("[%v] -> ", cursor.Value)
		cursor = cursor.next
	}
	result += fmt.Sprintf("[T]")

	return result
}

func (node *Node) Prev() *Node {
	return node.prev
}

func (node *Node) Next() *Node {
	return node.next
}
