package goclub_algo_linked_list

import (
	"fmt"
	"strconv"
)
// 最容易实现的单链表
type SingleListEasy struct {
	head *SingleListNode
}
func NewSingleListEasy(head *SingleListNode) *SingleListEasy {
	return &SingleListEasy{
		head: head,
	}
}

func (list *SingleListEasy) Head() (head *SingleListNode, hasHead bool) {
	if list.head == nil {
		return nil, false
	}
	return list.head, true
}
func (list *SingleListEasy) Tail() (tail *SingleListNode, hasTail bool) {
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
func (list *SingleListEasy) Length() (length int) {
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
func (list *SingleListEasy) Dump() (s string) {
	if list.head == nil {
		return ""
	}
	curr := list.head
	seen := map[*SingleListNode]int{}
	i := 0
	for {
		if index, ok := seen[curr]; ok {
			s = s + fmt.Sprintf("cycle{index:"+strconv.Itoa(index) + ", value:%s}", curr.value)
			return
		}
		s += fmt.Sprintf("%s", curr.value)
		if curr.next == nil{
			return
		}
		s += "->"
		seen[curr] = i
		i++
		curr = curr.next
	}
}
func (list *SingleListEasy) LeftPush(v interface{}) (newNode *SingleListNode) {
	newNode = NewSingleListNode(v)
	if list.head == nil {
		list.head = newNode
		return
	}
	newNode.next = list.head
	list.head = newNode
	return
}

func (list *SingleListEasy) RightPush(v interface{}) (newNode *SingleListNode) {
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
func (list *SingleListEasy) PrevNode(pivot *SingleListNode) (prevNode *SingleListNode, hasPrevNode bool) {
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

func (list *SingleListEasy) DeleteByNode(removeNode *SingleListNode) (ok bool) {
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
func (list *SingleListEasy) InsertBefore(pivot *SingleListNode, newNode *SingleListNode) (ok bool) {
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
func (list *SingleListEasy) InsertAfter(pivot *SingleListNode, newNode *SingleListNode) (ok bool) {
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
func (list *SingleListEasy) FindByIndex(index int) (target *SingleListNode, hasNode bool) {
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
func (list *SingleListEasy) Reverse() {
	// empty
	if list.head == nil {
		return
	}
	// one node 
	if list.head.next == nil {
		return
	}
	var prev *SingleListNode
	curr := list.head
	for {
		next := curr.next
		curr.next = prev
		if next == nil {
			list.head = curr
			return
		}
		prev = curr
		curr = next
	}
}

func (list *SingleListEasy) IsCycle() bool {
	if list.head == nil {
		return false
	}
	if list.head.next == nil {
		return false
	}
	slow , fast := list.head, list.head.next
	for {
		if fast == slow {
			return true
		}
		if fast == nil {
			return false
		}
		if fast.next == nil {
			return false
		}
		slow = slow.next
		fast = fast.next.next
	}
}