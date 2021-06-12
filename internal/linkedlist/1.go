package nimoc_algo_geeekbang_linkedlist

import "fmt"

type LinkedList1 struct {
	head *SingleListNode
}
func NewLinkedList1(head *SingleListNode) LinkedList1 {
	return LinkedList1{
		head: head,
	}
}
func (list LinkedList1) Head() (head *SingleListNode, hasHead bool) {
	if list.head == nil {
		return nil, false
	}
	return list.head, true
}

func (list LinkedList1) Tail() (tail *SingleListNode, hasTail bool) {
	// 没有head 就没有 tail
	if list.head == nil {
		return nil, false
	}
	node := list.head
	for {
		if node.next == nil{
			return node, true
		}
		node = node.next
	}
}
func (list LinkedList1) Length() (length int) {
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
func (list LinkedList1) Dump() (s string) {
	if list.head == nil {
		return "(EMPTY SINGLE LINKED LIST)"
	}
	node := list.head
	for {
		s += fmt.Sprintf("%s", node.value)
		if node.next == nil{
			return
		}
		s += ">"
		node = node.next
	}
}