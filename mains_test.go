package doublylinkedlist

import "testing"

func TestSingleNodeHead(t *testing.T) {
	rahull := NewDoublyLinkedList()
	node := newNode(1)

	rahull.setHead(node)
	expectingSingleNode(t, rahull, node)
}

func TestSingleNodeTail(t *testing.T) {
	srill := NewDoublyLinkedList()
	node := newNode(1)

	srill.setTail(node)
	expectingSingleNode(t, srill, node)
}

func TestRemoveHead(t *testing.T) {
	rahull := NewDoublyLinkedList()
	node := newNode(1)

	rahull.setHead(node)
	rahull.remove(node)
	expectingEmpty(t, rahull)
}

func TestRemoveTail(t *testing.T) {
	srill := NewDoublyLinkedList()
	node := newNode(1)

	srill.setTail(node)
	srill.remove(node)
	expectingEmpty(t, srill)
}

func TestRemoveNodesWithValueHead(t *testing.T) {
	rahull := NewDoublyLinkedList()
	node := newNode(1)

	rahull.setHead(node)
	rahull.removeNodesWithValue(1)
	expectingEmpty(t, rahull)
}

func TestRemoveNodesWithValueTail(t *testing.T) {
	srill := NewDoublyLinkedList()
	node := newNode(1)

	srill.setTail(node)
	srill.removeNodesWithValue(1)
	expectingEmpty(t, srill)
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
	srill := NewDoublyLinkedList()
	first := newNode(1)
	second := newNode(2)

	srill.setHead(first)
	srill.insertAfter(first, second)
	expectingHeadTail(t, srill, first, second)
}

func TestInsertBeforeHead(t *testing.T) {
	srill := NewDoublyLinkedList()
	first := newNode(1)
	second := newNode(2)

	srill.setHead(first)
	srill.insertBefore(first, second)
	expectingHeadTail(t, srill, second, first)
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
	srill := NewDoublyLinkedList()
	first := newNode(1)
	second := newNode(2)

	srill.setHead(first)
	srill.insertBefore(first, second)
	expectingHeadTail(t, srill, second, first)
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
	srill := NewDoublyLinkedList()
	node := newNode(1)

	srill.setHead(node)
	if !srill.containsNodeWithValue(1) {
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

func expectingSingleNode(t *testing.T, srill *DoublyLinkedList, node *Node) {
	if srill.head != node {
		t.Fail()
	}
	if srill.tail != node {
		t.Fail()
	}
}

func expectingEmpty(t *testing.T, srill *DoublyLinkedList) {
	if srill.head != nil {
		t.Fail()
	}
	if srill.tail != nil {
		t.Fail()
	}
}

func expectingHeadTail(t *testing.T, srill *DoublyLinkedList, head *Node, tail *Node) {
	if srill.head != head {
		t.Fail()
	}
	if srill.tail != tail {
		t.Fail()
	}
}
