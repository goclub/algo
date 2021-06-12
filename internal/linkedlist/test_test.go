package nimoc_algo_geeekbang_linkedlist

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewLinkedList1(t *testing.T) {
	{
		list := NewLinkedList1(nil)
		testSingleListNode(t, list)
	}
}
func testSingleListNode(t *testing.T, list SingleLinkedLister) {
	{
		head, hasHead := list.Head()
		assert.Nil(t, head)
		assert.Equal(t, hasHead, false)
	}
	{
		tail, hasTail := list.Tail()
		assert.Nil(t, tail)
		assert.Equal(t, hasTail, false)
	}
	{
		assert.Equal(t, list.Length(), 0)
	}
	{
		assert.Equal(t, list.Dump(), "(EMPTY SINGLE LINKED LIST)")
	}
	// {
	// 	list := newList()
	// 	{
	// 		head, hasHead := list.Head()
	// 		assert.Equal(t, head, (*SingleListNode)(nil))
	// 		assert.Equal(t, hasHead, false)
	// 		tail, hasTail := list.Tail()
	// 		assert.Equal(t, tail, (*SingleListNode)(nil))
	// 		assert.Equal(t, hasTail, false)
	// 		assert.Equal(t, list.Length(), uint(0))
	// 		assert.Equal(t, list.Dump(), "(EMPTY SINGLE LINKED LIST)")
	// 	}
	// 	list.RightPush("a")
	// 	{
	// 		head, hasHead := list.Head()
	// 		assert.Equal(t, head.value, "a")
	// 		assert.Equal(t, hasHead, true)
	// 		tail, hasTail := list.Tail()
	// 		assert.Equal(t, tail.value, "a")
	// 		assert.Equal(t, hasTail, true)
	// 		assert.Equal(t, list.Length(), uint(1))
	// 		assert.Equal(t, list.Dump(), "a")
	// 	}
	// 	list.RightPush("b")
	// 	{
	// 		head, hasHead := list.Head()
	// 		assert.Equal(t, head.value, "a")
	// 		assert.Equal(t, hasHead, true)
	// 		tail, hasTail := list.Tail()
	// 		assert.Equal(t, tail.value, "b")
	// 		assert.Equal(t, hasTail, true)
	// 		assert.Equal(t, list.Length(), uint(2))
	// 		assert.Equal(t, list.Dump(), "a->b")
	// 	}
	// 	list.RightPush("c")
	// 	{
	// 		head, hasHead := list.Head()
	// 		assert.Equal(t, head.value, "a")
	// 		assert.Equal(t, hasHead, true)
	// 		tail, hasTail := list.Tail()
	// 		assert.Equal(t, tail.value, "c")
	// 		assert.Equal(t, hasTail, true)
	// 		assert.Equal(t, list.Length(), uint(3))
	// 		assert.Equal(t, list.Dump(), "a->b->c")
	// 	}
	// 	list.RightPush("d")
	// 	{
	// 		head, hasHead := list.Head()
	// 		assert.Equal(t, head.value, "a")
	// 		assert.Equal(t, hasHead, true)
	// 		tail, hasTail := list.Tail()
	// 		assert.Equal(t, tail.value, "d")
	// 		assert.Equal(t, hasTail, true)
	// 		assert.Equal(t, list.Length(), uint(4))
	// 		assert.Equal(t, list.Dump(), "a->b->c->d")
	// 	}
	// }
	// // list.LeftPush() list.Head() list.Tail() list.Length()
	// {
	// 	list := newList()
	// 	{
	// 		head, hasHead := list.Head()
	// 		assert.Equal(t, head, (*SingleListNode)(nil))
	// 		assert.Equal(t, hasHead, false)
	// 		tail, hasTail := list.Tail()
	// 		assert.Equal(t, tail, (*SingleListNode)(nil))
	// 		assert.Equal(t, hasTail, false)
	// 		assert.Equal(t, list.Length(), uint(0))
	// 		assert.Equal(t, list.Dump(), "(EMPTY SINGLE LINKED LIST)")
	// 	}
	// 	list.LeftPush("a")
	// 	{
	// 		head, hasHead := list.Head()
	// 		assert.Equal(t, head.value, "a")
	// 		assert.Equal(t, hasHead, true)
	// 		tail, hasTail := list.Tail()
	// 		assert.Equal(t, tail.value, "a")
	// 		assert.Equal(t, hasTail, true)
	// 		assert.Equal(t, list.Length(), uint(1))
	// 		assert.Equal(t, list.Dump(), "a")
	// 	}
	// 	list.LeftPush("b")
	// 	{
	// 		head, hasHead := list.Head()
	// 		assert.Equal(t, head.value, "b")
	// 		assert.Equal(t, hasHead, true)
	// 		tail, hasTail := list.Tail()
	// 		assert.Equal(t, tail.value, "a")
	// 		assert.Equal(t, hasTail, true)
	// 		assert.Equal(t, list.Length(), uint(2))
	// 		assert.Equal(t, list.Dump(), "b->a")
	// 	}
	// 	list.LeftPush("c")
	// 	{
	// 		head, hasHead := list.Head()
	// 		assert.Equal(t, head.value, "c")
	// 		assert.Equal(t, hasHead, true)
	// 		tail, hasTail := list.Tail()
	// 		assert.Equal(t, tail.value, "a")
	// 		assert.Equal(t, hasTail, true)
	// 		assert.Equal(t, list.Length(), uint(3))
	// 		assert.Equal(t, list.Dump(), "c->b->a")
	// 	}
	// }
	// // node.Next() node.Value()
	// {
	// 	{
	// 		node := NewSingleListNode("a")
	// 		nextNode, hasNext := node.Next()
	// 		assert.Nil(t, nextNode)
	// 		assert.Equal(t, hasNext, false)
	// 		assert.Equal(t, node.value, node.Value())
	// 		assert.Equal(t, node.value, "a")
	// 	}
	// 	{
	// 		node := NewSingleListNode("a")
	// 		node.next = NewSingleListNode("b")
	// 		nextNode, hasNext := node.Next()
	// 		assert.Equal(t, nextNode.Value(), "b")
	// 		assert.Equal(t, hasNext, true)
	// 		assert.Equal(t, node.value, node.Value())
	// 		assert.Equal(t, node.value, "a")
	// 	}
	// }
	// // // list.FindByIndex
	// {
	// 	list := newList()
	// 	{
	// 		node, hasNode := list.FindByIndex(0)
	// 		assert.Equal(t, hasNode, false)
	// 		assert.Equal(t, node, (*SingleListNode)(nil))
	// 	}
	// 	list.RightPush("a")
	// 	{
	// 		node, hasNode := list.FindByIndex(0)
	// 		assert.Equal(t, hasNode, true)
	// 		assert.Equal(t, node.Value(), "a")
	// 	}
	// 	{
	// 		node, hasNode := list.FindByIndex(1)
	// 		assert.Equal(t, hasNode, false)
	// 		assert.Equal(t, node, (*SingleListNode)(nil))
	// 	}
	// 	list.RightPush("b")
	// 	{
	// 		node, hasNode := list.FindByIndex(0)
	// 		assert.Equal(t, hasNode, true)
	// 		assert.Equal(t, node.Value(), "a")
	// 	}
	// 	{
	// 		node, hasNode := list.FindByIndex(1)
	// 		assert.Equal(t, hasNode, true)
	// 		assert.Equal(t, node.Value(), "b")
	// 	}
	// 	{
	// 		node, hasNode := list.FindByIndex(2)
	// 		assert.Equal(t, hasNode, false)
	// 		assert.Equal(t, node, (*SingleListNode)(nil))
	// 	}
	// 	list.RightPush("c")
	// 	{
	// 		node, hasNode := list.FindByIndex(0)
	// 		assert.Equal(t, hasNode, true)
	// 		assert.Equal(t, node.Value(), "a")
	// 	}
	// 	{
	// 		node, hasNode := list.FindByIndex(1)
	// 		assert.Equal(t, hasNode, true)
	// 		assert.Equal(t, node.Value(), "b")
	// 	}
	// 	{
	// 		node, hasNode := list.FindByIndex(2)
	// 		assert.Equal(t, hasNode, true)
	// 		assert.Equal(t, node.Value(), "c")
	// 	}
	// 	{
	// 		node, hasNode := list.FindByIndex(3)
	// 		assert.Equal(t, hasNode, false)
	// 		assert.Equal(t, node, (*SingleListNode)(nil))
	// 	}
	// }
	// // // list.PrevNode
	// {
	// 	list := newList()
	// 	{
	// 		prevNode, hasPrevNode := list.PrevNode(NewSingleListNode("test length 0"))
	// 		assert.Equal(t, hasPrevNode, false)
	// 		assert.Equal(t, prevNode, (*SingleListNode)(nil))
	// 	}
	// 	list.RightPush("a")
	// 	list.RightPush("b")
	// 	list.RightPush("c")
	// 	assert.Equal(t, list.Dump(), "a->b->c")
	// 	head, hasHead := list.Head()
	// 	assert.Equal(t, hasHead, true)
	// 	{
	// 		prevNode, hasPrevNode := list.PrevNode(head)
	// 		assert.Equal(t, hasPrevNode, false)
	// 		assert.Equal(t, prevNode, (*SingleListNode)(nil))
	// 	}
	// 	{
	// 		prevNode, hasPrevNode := list.PrevNode(head.next)
	// 		assert.Equal(t, hasPrevNode, true)
	// 		assert.Equal(t, prevNode.value, "a")
	// 	}
	// 	{
	// 		prevNode, hasPrevNode := list.PrevNode(head.next.next)
	// 		assert.Equal(t, hasPrevNode, true)
	// 		assert.Equal(t, prevNode.value, "b")
	// 	}
	// 	{
	// 		prevNode, hasPrevNode := list.PrevNode(head.next.next.next)
	// 		assert.Equal(t, hasPrevNode, false)
	// 		assert.Equal(t, prevNode, (*SingleListNode)(nil))
	// 	}
	// 	{
	// 		prevNode, hasPrevNode := list.PrevNode(NewSingleListNode("not exist"))
	// 		assert.Equal(t, hasPrevNode, false)
	// 		assert.Equal(t, prevNode, (*SingleListNode)(nil))
	// 	}
	// }
	// // list.DeleteByNode()
	// {
	// 	// EMPTY
	// 	{
	// 		list := newList()
	// 		assert.Equal(t, list.DeleteByNode(nil), false)
	// 	}
	// 	// HEAD_NODE ONE_NODE
	// 	{
	// 		list := newList()
	// 		list.RightPush("a")
	// 		node, hasNode := list.Head()
	// 		assert.Equal(t, hasNode, true)
	// 		assert.Equal(t, list.DeleteByNode(node), true)
	// 		assert.Equal(t, list.Length(), uint(0))
	// 	}
	// 	// HEAD_NODE TWO_NODE
	// 	{
	// 		list := newList()
	// 		list.RightPush("a")
	// 		list.RightPush("b")
	// 		node, hasNode := list.Head()
	// 		assert.Equal(t, hasNode, true)
	// 		assert.Equal(t, list.DeleteByNode(node), true)
	// 		assert.Equal(t, list.Length(), uint(1))
	// 		assert.Equal(t, list.Dump(), "b")
	// 	}
	// 	// HEAD_NODE THREE_NODE
	// 	{
	// 		list := newList()
	// 		list.RightPush("a")
	// 		list.RightPush("b")
	// 		list.RightPush("c")
	// 		node, hasNode := list.Head()
	// 		assert.Equal(t, hasNode, true)
	// 		assert.Equal(t, list.DeleteByNode(node), true)
	// 		assert.Equal(t, list.Length(), uint(2))
	// 		assert.Equal(t, list.Dump(), "b->c")
	// 	}
	// 	// TAIL_NODE ONE_NODE
	// 	{
	// 		list := newList()
	// 		list.RightPush("a")
	// 		node, hasNode := list.Tail()
	// 		assert.Equal(t, hasNode, true)
	// 		assert.Equal(t, list.DeleteByNode(node), true)
	// 		assert.Equal(t, list.Length(), uint(0))
	// 		assert.Equal(t, list.Dump(), "(EMPTY SINGLE LINKED LIST)")
	// 	}
	// 	// TAIL_NODE TWO_NODE
	// 	{
	// 		list := newList()
	// 		list.RightPush("a")
	// 		list.RightPush("b")
	// 		node, hasNode := list.Tail()
	// 		assert.Equal(t, hasNode, true)
	// 		assert.Equal(t, list.DeleteByNode(node), true)
	// 		assert.Equal(t, list.Length(), uint(1))
	// 		assert.Equal(t, list.Dump(), "a")
	// 	}
	// }
}
