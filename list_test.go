package doublylinkedlist

import (
	"fmt"
	"testing"
)

func TestList_Len(t *testing.T) {
	list := NewList()
	list.PushFront(7)
	list.PushFront(3)
	list.PushBack(29)

	if result := list.Len(); result != 3 {
		t.Fatalf("Wrong total quantity of nodes (length): expected 3, got %d", result)
	}
}

func TestList_FirstLast(t *testing.T) {
	list := NewList()
	list.PushFront(7)
	firstNode := list.PushFront(3)
	lastNode := list.PushBack(9)

	if list.First() != firstNode {
		t.Fatalf("Wrong first node in list")
	}

	if list.Last() != lastNode {
		t.Fatalf("Wrong last node in list")
	}
}

func TestList_PushFront(t *testing.T) {
	expected := 5
	list := NewList()
	list.PushFront(7)
	list.PushFront(expected)

	if result := list.First().Value; result != expected {
		t.Fatalf("Wrong value of first node: expected %d, got %d", expected, result)
	}
}

func TestList_PushBack(t *testing.T) {
	expected := 12
	list := NewList()
	list.PushBack(10)
	list.PushBack(expected)

	if result := list.Last().Value; result != expected {
		t.Fatalf("Wrong value of last node: expected %d, got %d", expected, result)
	}
}

func TestList_PushBefore(t *testing.T) {
	list := NewList()
	node := list.PushBack(5)
	list.PushBack(9)
	node = list.PushBefore(node, 3)

	if list.First() != node {
		fmt.Printf("Result: %s\n", list)
		t.Fatalf("Wrong node pushed before other node")
	}
}

func TestList_PushAfter(t *testing.T) {
	list := NewList()
	list.PushFront(1)
	node := list.PushBack(5)
	node = list.PushAfter(node, 3)

	if list.Last() != node {
		t.Fatalf("Wrong node pushed after other node")
	}
}

func TestList_Remove(t *testing.T) {
	list := NewList()
	list.PushFront(2)
	list.PushBack(19)
	firstNode := list.PushFront(1)
	lastNode := list.PushBack(27)
	list.PushAfter(firstNode, 5)
	list.PushBefore(lastNode, 8)

	list.Remove(firstNode)
	if list.Len() != 5 || list.First() == firstNode {
		t.Fatalf("Failed to remove first node")
	}

	list.Remove(lastNode)
	if list.Len() != 4 || list.Last() == lastNode {
		t.Fatalf("Failed to remove last node")
	}
}
