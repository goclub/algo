package nimoc_algo_geeekbang_linkedlist

import (
	"fmt"
)

type LinkedList1 struct {
	head *SingleListNode
}
func NewLinkedList1(head *SingleListNode) *LinkedList1 {
	return &LinkedList1{
		head: head,
	}
}

func (list *LinkedList1) Head() (head *SingleListNode, hasHead bool) {
	if list.head == nil {
		return nil, false
	}
	return list.head, true
}
func (list *LinkedList1) Tail() (tail *SingleListNode, hasTail bool) {
	if list.head == nil {
		return nil, false
	}
	node := list.head
	for {
		if node.next == nil {
			return node, true
		}
		node = node.next
	}
}
func (list *LinkedList1) Length() (length int) {
	if list.head == nil {
		return 0
	}
	node := list.head
	for {
		length++
		if node.next == nil{
			return
		}
		node = node.next
	}
}
func (list *LinkedList1) Dump() (s string) {
	if list.head == nil {
		return ""
	}
	node := list.head
	for {
		s += fmt.Sprintf("%s", node.value)
		if node.next == nil{
			return
		}
		s += "->"
		node = node.next
	}
}
func (list *LinkedList1) LeftPush(v interface{}) (newNode *SingleListNode) {
	newNode = NewSingleListNode(v)
	if list.head == nil {
		list.head = newNode
		return
	}
	newNode.next = list.head
	list.head = newNode
	return
}

func (list *LinkedList1) RightPush(v interface{}) (newNode *SingleListNode) {
	newNode = NewSingleListNode(v)
	if list.head == nil {
		list.head = newNode
		return
	}
	node := list.head
	for {
		if node.next == nil {
			node.next = newNode
			return
		}
		node = node.next
	}
}
func (list *LinkedList1) PrevNode(pivot *SingleListNode) (prevNode *SingleListNode, hasPrevNode bool) {
	if list.head == nil {
		return nil, false
	}
	node := list.head
	for {
		if node.next == pivot {
			return node, true
		}
		if node.next == nil {
			return nil, false
		}
		node = node.next
	}
}

func (list *LinkedList1) DeleteByNode(removeNode *SingleListNode) (ok bool) {
	if list.head == nil {
		return false
	}
	if list.head == removeNode {
		list.head = list.head.next
		removeNode.next = nil
		return true
	}
	node := list.head

	for {
		if node.next == removeNode {
			node.next = node.next.next
			removeNode.next = nil
			return true
		}
		if node.next == nil {
			return false
		}
		node = node.next
	}
}
func (list *LinkedList1) InsertBefore(pivot *SingleListNode, newNode *SingleListNode) (ok bool) {
	if list.head == nil {
		return false
	}
	if list.head == pivot {
		newNode.next = list.head
		list.head = newNode
		return true
	}
	node := list.head
	for {
		if node.next == nil {
			return false
		}
		if node.next == pivot {
			newNode.next = node.next
			node.next = newNode
			return true
		}
		node = node.next
	}
}
func (list *LinkedList1) InsertAfter(pivot *SingleListNode, newNode *SingleListNode) (ok bool) {
	if list.head == nil {
		return false
	}
	node := list.head
	for {
		if node == pivot {
			newNode.next = node.next
			node.next = newNode
			return true
		}
		if node.next == nil {
			return false
		}
		node = node.next
	}
}
func (list *LinkedList1) FindByIndex(index int) (target *SingleListNode, hasNode bool) {
	if list.head == nil {
		return nil, false
	}
	node := list.head
	for count := 0; ;count++ {
		if count == index {
			return node, true
		}
		if node.next == nil {
			return nil, false
		}
		node = node.next
	}
}
func (list *LinkedList1) Reverse() {
	// empty
	if list.head == nil {
		return
	}
	// one node
	if list.head.next == nil {
		return
	}
	// eg: a->b->c->d->e
	var prev *SingleListNode
	for curr := list.head; curr != nil;{
		// I: curr is a
		// 	  prev is nil
		//    next is b

		// II: curr is b
		// 	  prev is a
		//    next is c
		next := curr.next
		// I: nil<-a b->c->d->e

		// II: nil<-a<-b c->d->e
		curr.next = prev
		prev = curr
		curr = next
	}
	list.head = prev
}