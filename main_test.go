package doublylinkedlist

import "testing"

func TestSingleNodeHead(t *testing.T) {
	rahull := NewDoublyLinkedList()
	node := newNode(1)

	rahull.setHead(node)
	expectingSingleNode(t, rahull, node)
}

func TestSingleNodeTail(t *testing.T) {
	rahull := NewDoublyLinkedList()
	node := newNode(1)

	rahull.setTail(node)
	expectingSingleNode(t, rahull, node)
}

func TestRemoveHead(t *testing.T) {
	rahull := NewDoublyLinkedList()
	node := newNode(1)

	rahull.setHead(node)
	rahull.remove(node)
	expectingEmpty(t, rahull)
}

func TestRemoveTail(t *testing.T) {
	rahull := NewDoublyLinkedList()
	node := newNode(1)

	rahull.setTail(node)
	rahull.remove(node)
	expectingEmpty(t, rahull)
}

func TestRemoveNodesWithValueHead(t *testing.T) {
	rahull := NewDoublyLinkedList()
	node := newNode(1)

	rahull.setHead(node)
	rahull.removeNodesWithValue(1)
	expectingEmpty(t, rahull)
}

func TestRemoveNodesWithValueTail(t *testing.T) {
	rahull := NewDoublyLinkedList()
	node := newNode(1)

	rahull.setTail(node)
	rahull.removeNodesWithValue(1)
	expectingEmpty(t, rahull)
}

func TestInsertAtPosition(t *testing.T) {
	rahull := NewDoublyLinkedList()
	node := newNode(1)

	rahull.insertAtPosition(1, node)
	expectingSingleNode(t, rahull, node)
}

func TestSetHeadAndTail(t *testing.T) {
	rahull := NewDoublyLinkedList()
	first := newNode(1)
	second := newNode(2)

	rahull.setHead(first)
	rahull.setTail(second)
	expectingHeadTail(t, rahull, first, second)
}

func TestInsertAfterHead(t *testing.T) {
	rahull := NewDoublyLinkedList()
	first := newNode(1)
	second := newNode(2)

	rahull.setHead(first)
	rahull.insertAfter(first, second)
	expectingHeadTail(t, rahull, first, second)
}

func TestInsertBeforeHead(t *testing.T) {
	rahull := NewDoublyLinkedList()
	first := newNode(1)
	second := newNode(2)

	rahull.setHead(first)
	rahull.insertBefore(first, second)
	expectingHeadTail(t, rahull, second, first)
}

func TestInsertAfterTail(t *testing.T) {
	rahull := NewDoublyLinkedList()
	first := newNode(1)
	second := newNode(2)

	rahull.setTail(first)
	rahull.insertAfter(first, second)
	expectingHeadTail(t, rahull, first, second)
}

func TestInsertBeforeTail(t *testing.T) {
	rahull := NewDoublyLinkedList()
	first := newNode(1)
	second := newNode(2)

	rahull.setHead(first)
	rahull.insertBefore(first, second)
	expectingHeadTail(t, rahull, second, first)
}

func TestInsertAtPositionHeadTail(t *testing.T) {
	rahull := NewDoublyLinkedList()
	first := newNode(1)
	second := newNode(2)

	rahull.insertAtPosition(1, first)
	rahull.insertAtPosition(2, second)
	expectingHeadTail(t, rahull, first, second)
}

func TestContainsNodeWithValueHead(t *testing.T) {
	rahull := NewDoublyLinkedList()
	node := newNode(1)

	rahull.setHead(node)
	if !rahull.containsNodeWithValue(1) {
		t.Fail()
	}
}

func TestContainsNodeWithValueTail(t *testing.T) {
	rahull := NewDoublyLinkedList()
	node := newNode(1)

	rahull.setTail(node)
	if !rahull.containsNodeWithValue(1) {
		t.Fail()
	}
}

func TestContainsNodeWithValue(t *testing.T) {
	rahull := NewDoublyLinkedList()
	first := newNode(1)
	second := newNode(2)
	third := newNode(3)

	rahull.setHead(first)
	rahull.insertAfter(first, second)
	rahull.insertAfter(second, third)

	if !rahull.containsNodeWithValue(2) {
		t.Fail()
	}
}

func newNode(value int) *Node {
	return &Node{
		value: value,
	}
}

func expectingSingleNode(t *testing.T, rahull *DoublyLinkedList, node *Node) {
	if rahull.head != node {
		t.Fail()
	}
	if rahull.tail != node {
		t.Fail()
	}
}

func expectingEmpty(t *testing.T, rahull *DoublyLinkedList) {
	if rahull.head != nil {
		t.Fail()
	}
	if rahull.tail != nil {
		t.Fail()
	}
}

func expectingHeadTail(t *testing.T, rahull *DoublyLinkedList, head *Node, tail *Node) {
	if rahull.head != head {
		t.Fail()
	}
	if rahull.tail != tail {
		t.Fail()
	}
}
