package goclub_algo_linked_list

import (
	"fmt"
	"strconv"
)
// 最容易实现的单链表
type SingleListEasy struct {
	first *SingleListNode
}
func NewSingleListEasy(first *SingleListNode) *SingleListEasy {
	return &SingleListEasy{
		first: first,
	}
}

func (list *SingleListEasy) First() (first *SingleListNode, hasFirst bool) {
	if list.first == nil {
		return nil, false
	}
	return list.first, true
}
func (list *SingleListEasy) Last() (last *SingleListNode, hasLast bool) {
	if list.first == nil {
		return nil, false
	}
	node := list.first
	for {
		if node.next == nil {
			return node, true
		}
		node = node.next
	}
}
func (list *SingleListEasy) Size() (length int) {
	if list.first == nil {
		return 0
	}
	node := list.first
	for {
		length++
		if node.next == nil{
			return
		}
		node = node.next
	}
}
func (list *SingleListEasy) Dump() (s string) {
	if list.first == nil {
		return ""
	}
	curr := list.first
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
func (list *SingleListEasy) AddFirst(v interface{}) (newNode *SingleListNode) {
	newNode = NewSingleListNode(v)
	if list.first == nil {
		list.first = newNode
		return
	}
	newNode.next = list.first
	list.first = newNode
	return
}

func (list *SingleListEasy) AddLast(v interface{}) (newNode *SingleListNode) {
	newNode = NewSingleListNode(v)
	if list.first == nil {
		list.first = newNode
		return
	}
	node := list.first
	for {
		if node.next == nil {
			node.next = newNode
			return
		}
		node = node.next
	}
}
func (list *SingleListEasy) PrevNode(pivot *SingleListNode) (prevNode *SingleListNode, hasPrevNode bool) {
	if list.first == nil {
		return nil, false
	}
	node := list.first
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

func (list *SingleListEasy) RemoveByNode(removeNode *SingleListNode) (ok bool) {
	if list.first == nil {
		return false
	}
	if list.first == removeNode {
		list.first = list.first.next
		removeNode.next = nil
		return true
	}
	node := list.first

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
	if list.first == nil {
		return false
	}
	if list.first == pivot {
		newNode.next = list.first
		list.first = newNode
		return true
	}
	node := list.first
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
	if list.first == nil {
		return false
	}
	node := list.first
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
	if list.first == nil {
		return nil, false
	}
	node := list.first
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
	if list.first == nil {
		return
	}
	// one node 
	if list.first.next == nil {
		return
	}
	var prev *SingleListNode
	curr := list.first
	for {
		next := curr.next
		curr.next = prev
		if next == nil {
			list.first = curr
			return
		}
		prev = curr
		curr = next
	}
}

func (list *SingleListEasy) IsCycle() bool {
	if list.first == nil {
		return false
	}
	if list.first.next == nil {
		return false
	}
	slow , fast := list.first, list.first.next
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