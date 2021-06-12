package nimoc_algo_geeekbang_linkedlist

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewLinkedList1(t *testing.T) {
	{
		testSingleListNode(t, func(node *SingleListNode) SingleLinkedLister {
			return NewLinkedList1(node)
		})
	}
}
type equalData struct {
	HeadValue interface{}
	HasHead bool
	TailValue interface{}
	HasTail bool
	Length int
	Dump string
}
func equalList(t *testing.T, list SingleLinkedLister, data equalData) {
	// Dump
	assert.Equal(t, data.Dump, list.Dump())
	// head
	{
		head, hasHead := list.Head()
		assert.Equal(t, data.HasHead, hasHead)
		if hasHead {
			assert.Equal(t, data.HeadValue, head.Value())
		} else {
			assert.Equal(t, data.HeadValue, "")
		}
	}
	// tail
	{
		tail, hasTail := list.Tail()
		assert.Equal(t, data.HasTail, hasTail)
		if hasTail {
			assert.Equal(t, data.TailValue, tail.Value())
		} else {
			assert.Equal(t, data.TailValue, "")
		}
	}
	// length
	assert.Equal(t, data.Length, list.Length())
}
func testSingleListNode(t *testing.T, newList func(node *SingleListNode) SingleLinkedLister ) {
	{
		list := newList(nil)
		equalList(t, list, equalData{
			HeadValue:    "",
			HasHead: false,
			TailValue:    "",
			HasTail: false,
			Length:  0,
			Dump:    "",
		})
	}
	// LeftPush
	// a
	// b->a
	// c->b->a
	{
		list := newList(NewSingleListNode("a"))
		equalList(t, list, equalData{
			HeadValue:    "a",
			HasHead: true,
			TailValue:    "a",
			HasTail: true,
			Length:  1,
			Dump:    "a",
		})
		list.LeftPush("b")
		equalList(t, list, equalData{
			Dump:    "b->a",
			HeadValue:    "b",
			HasHead: true,
			TailValue:    "a",
			HasTail: true,
			Length:  2,
		})
		list.LeftPush("c")
		equalList(t, list, equalData{
			Dump:    "c->b->a",
			HeadValue:    "c",
			HasHead: true,
			TailValue:    "a",
			HasTail: true,
			Length:  3,
		})
	}
	// RightPush
	// a
	// a->b
	// a->b->c
	{
		list := newList(NewSingleListNode("a"))
		equalList(t, list, equalData{
			HeadValue:    "a",
			HasHead: true,
			TailValue:    "a",
			HasTail: true,
			Length:  1,
			Dump:    "a",
		})
		list.RightPush("b")
		equalList(t, list, equalData{
			Dump:    "a->b",
			HeadValue:    "a",
			HasHead: true,
			TailValue:    "b",
			HasTail: true,
			Length:  2,
		})
		list.RightPush("c")
		equalList(t, list, equalData{
			Dump:    "a->b->c",
			HeadValue:    "a",
			HasHead: true,
			TailValue:    "c",
			HasTail: true,
			Length:  3,
		})
	}
	// PrevNode
	{
		list := newList(nil)
		a := list.RightPush("a")
		b := list.RightPush("b")
		c := list.RightPush("c")
		equalList(t, list, equalData{
			Dump:    "a->b->c",
			HeadValue:    "a",
			HasHead: true,
			TailValue:    "c",
			HasTail: true,
			Length:  3,
		})
		// head
		{
			prevNode, hasPrevNode := list.PrevNode(a)
			assert.Nil(t, prevNode)
			assert.Equal(t, hasPrevNode, false)
		}
		// middle
		{
			prevNode, hasPrevNode := list.PrevNode(b)
			assert.Equal(t, a, prevNode)
			assert.Equal(t, hasPrevNode, true)
		}
		// tail
		{
			prevNode, hasPrevNode := list.PrevNode(c)
			assert.Equal(t, b, prevNode)
			assert.Equal(t, hasPrevNode, true)
		}
		// not found node
		{
			prevNode, hasPrevNode := list.PrevNode(NewSingleListNode(""))
			assert.Equal(t, hasPrevNode, false)
			assert.Nil(t, prevNode)
		}
	}
	// DeleteByNode
	{
		list := newList(nil)
		a := list.RightPush("a")
		b := list.RightPush("b")
		c := list.RightPush("c")
		equalList(t, list, equalData{
			Dump:    "a->b->c",
			HeadValue:    "a",
			HasHead: true,
			TailValue:    "c",
			HasTail: true,
			Length:  3,
		})
		// head
		{
			assert.Equal(t, list.DeleteByNode(a), true)
			equalList(t, list, equalData{
				Dump:    "b->c",
				HeadValue:    "b",
				HasHead: true,
				TailValue:    "c",
				HasTail: true,
				Length:  2,
			})
			assert.Equal(t, list.DeleteByNode(a), false)
		}
		// tail
		{
			assert.Equal(t, list.DeleteByNode(c), true)
			equalList(t, list, equalData{
				Dump:    "b",
				HeadValue:    "b",
				HasHead: true,
				TailValue:    "b",
				HasTail: true,
				Length:  1,
			})
		}
		// one node
		{
			assert.Equal(t, list.DeleteByNode(b), true)
			equalList(t, list, equalData{
				Dump:    "",
				HeadValue:    "",
				HasHead: false,
				TailValue:    "",
				HasTail: false,
				Length:  0,
			})
		}
		// middle
		{
			_ = list.RightPush("x")
			y := list.RightPush("y")
			_ = list.RightPush("z")
			equalList(t, list, equalData{
				Dump:    "x->y->z",
				HeadValue:    "x",
				HasHead: true,
				TailValue:    "z",
				HasTail: true,
				Length:  3,
			})
			assert.Equal(t, list.DeleteByNode(y), true)
			equalList(t, list, equalData{
				Dump:    "x->z",
				HeadValue:    "x",
				HasHead: true,
				TailValue:    "z",
				HasTail: true,
				Length:  2,
			})
		}
	}
	// InsertBefore
	{
		list := newList(nil)
		a := list.RightPush("a")
		// head
		b := NewSingleListNode("b")
		assert.Equal(t, list.InsertBefore(a, b), true)
		equalList(t, list, equalData{
			Dump:    "b->a",
			HeadValue:    "b",
			HasHead: true,
			TailValue:    "a",
			HasTail: true,
			Length:  2,
		})
		c := NewSingleListNode("c")
		assert.Equal(t, list.InsertBefore(b, c), true)
		equalList(t, list, equalData{
			Dump:    "c->b->a",
			HeadValue:    "c",
			HasHead: true,
			TailValue:    "a",
			HasTail: true,
			Length:  3,
		})
		// middle
		x := NewSingleListNode("x")
		assert.Equal(t, list.InsertBefore(b, x), true)
		equalList(t, list, equalData{
			Dump:    "c->x->b->a",
			HeadValue:    "c",
			HasHead: true,
			TailValue:    "a",
			HasTail: true,
			Length:  4,
		})
		// tail
		z := NewSingleListNode("z")
		assert.Equal(t, list.InsertBefore(a, z), true)
		equalList(t, list, equalData{
			Dump:    "c->x->b->z->a",
			HeadValue:    "c",
			HasHead: true,
			TailValue:    "a",
			HasTail: true,
			Length:  5,
		})

	}
}
