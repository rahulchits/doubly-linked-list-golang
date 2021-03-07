package doublylinkedlist

import "testing"

func TestSingleNodeHead(t *testing.T) {
	tll := NewDoublyLinkedList()
	node := newNode(1)

	tll.setHead(node)
	expectingSingleNode(t, tll, node)
}

func TestSingleNodeTail(t *testing.T) {
	tll := NewDoublyLinkedList()
	node := newNode(1)

	tll.setTail(node)
	expectingSingleNode(t, tll, node)
}

func TestRemoveHead(t *testing.T) {
	tll := NewDoublyLinkedList()
	node := newNode(1)

	tll.setHead(node)
	tll.remove(node)
	expectingEmpty(t, tll)
}

func TestRemoveTail(t *testing.T) {
	tll := NewDoublyLinkedList()
	node := newNode(1)

	tll.setTail(node)
	tll.remove(node)
	expectingEmpty(t, tll)
}

func TestRemoveNodesWithValueHead(t *testing.T) {
	tll := NewDoublyLinkedList()
	node := newNode(1)

	tll.setHead(node)
	tll.removeNodesWithValue(1)
	expectingEmpty(t, tll)
}

func TestRemoveNodesWithValueTail(t *testing.T) {
	tll := NewDoublyLinkedList()
	node := newNode(1)

	tll.setTail(node)
	tll.removeNodesWithValue(1)
	expectingEmpty(t, tll)
}

func TestInsertAtPosition(t *testing.T) {
	tll := NewDoublyLinkedList()
	node := newNode(1)

	tll.insertAtPosition(1, node)
	expectingSingleNode(t, tll, node)
}

func TestSetHeadAndTail(t *testing.T) {
	tll := NewDoublyLinkedList()
	first := newNode(1)
	second := newNode(2)

	tll.setHead(first)
	tll.setTail(second)
	expectingHeadTail(t, tll, first, second)
}

func TestInsertAfterHead(t *testing.T) {
	tll := NewDoublyLinkedList()
	first := newNode(1)
	second := newNode(2)

	tll.setHead(first)
	tll.insertAfter(first, second)
	expectingHeadTail(t, tll, first, second)
}

func TestInsertBeforeHead(t *testing.T) {
	tll := NewDoublyLinkedList()
	first := newNode(1)
	second := newNode(2)

	tll.setHead(first)
	tll.insertBefore(first, second)
	expectingHeadTail(t, tll, second, first)
}

func TestInsertAfterTail(t *testing.T) {
	tll := NewDoublyLinkedList()
	first := newNode(1)
	second := newNode(2)

	tll.setTail(first)
	tll.insertAfter(first, second)
	expectingHeadTail(t, tll, first, second)
}

func TestInsertBeforeTail(t *testing.T) {
	tll := NewDoublyLinkedList()
	first := newNode(1)
	second := newNode(2)

	tll.setHead(first)
	tll.insertBefore(first, second)
	expectingHeadTail(t, tll, second, first)
}

func TestInsertAtPositionHeadTail(t *testing.T) {
	tll := NewDoublyLinkedList()
	first := newNode(1)
	second := newNode(2)

	tll.insertAtPosition(1, first)
	tll.insertAtPosition(2, second)
	expectingHeadTail(t, tll, first, second)
}

func TestContainsNodeWithValueHead(t *testing.T) {
	tll := NewDoublyLinkedList()
	node := newNode(1)

	tll.setHead(node)
	if !tll.containsNodeWithValue(1) {
		t.Fail()
	}
}

func TestContainsNodeWithValueTail(t *testing.T) {
	tll := NewDoublyLinkedList()
	node := newNode(1)

	tll.setTail(node)
	if !tll.containsNodeWithValue(1) {
		t.Fail()
	}
}

func TestContainsNodeWithValue(t *testing.T) {
	tll := NewDoublyLinkedList()
	first := newNode(1)
	second := newNode(2)
	third := newNode(3)

	tll.setHead(first)
	tll.insertAfter(first, second)
	tll.insertAfter(second, third)

	if !tll.containsNodeWithValue(2) {
		t.Fail()
	}
}

func newNode(value int) *Node {
	return &Node{
		value: value,
	}
}

func expectingSingleNode(t *testing.T, tll *DoublyLinkedList, node *Node) {
	if tll.head != node {
		t.Fail()
	}
	if tll.tail != node {
		t.Fail()
	}
}

func expectingEmpty(t *testing.T, tll *DoublyLinkedList) {
	if tll.head != nil {
		t.Fail()
	}
	if tll.tail != nil {
		t.Fail()
	}
}

func expectingHeadTail(t *testing.T, tll *DoublyLinkedList, head *Node, tail *Node) {
	if tll.head != head {
		t.Fail()
	}
	if tll.tail != tail {
		t.Fail()
	}
}
