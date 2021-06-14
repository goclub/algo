package nimoc_algo_geeekbang_linkedlist

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewSingleList1(t *testing.T) {
	{
		testSingleListNode(t, func(node *SingleListNode) SingleLister {
			return NewSingleList1(node)
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
func equalList(t *testing.T, list SingleLister, data equalData) {
	// Dump
	assert.Equal(t, data.Dump, list.Dump())
	// head
	{
		head, hasHead := list.Head()
		assert.Equal(t, data.HasHead, hasHead)
		if hasHead {
			assert.Equal(t, data.HeadValue, head.Value())
		} else {
			assert.Equal(t, data.HeadValue, nil)
		}
	}
	// tail
	{
		tail, hasTail := list.Tail()
		assert.Equal(t, data.HasTail, hasTail)
		if hasTail {
			assert.Equal(t, data.TailValue, tail.Value())
		} else {
			assert.Equal(t, data.TailValue, nil)
		}
	}
	// length
	assert.Equal(t, data.Length, list.Length())
}
func testSingleListNode(t *testing.T, newList func(node *SingleListNode) SingleLister ) {
	{
		list := newList(nil)
		equalList(t, list, equalData{
			HeadValue:    nil,
			HasHead: false,
			TailValue:    nil,
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
		list := newList(nil)
		list.LeftPush("a")
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
	// a->b->c
	// PrevNode(a) return nil, false
	// PrevNode(b) return a, true
	// PrevNode(c) return b, true
	// PrevNode(x) return nil, false
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
		// a->b->c
		// delete a, list is b->c
		// delete a, return false
		// delete c, list is b
		// delete b, list is
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
				HeadValue:    nil,
				HasHead: false,
				TailValue:    nil,
				HasTail: false,
				Length:  0,
			})
		}
		// middle
		// x->y->z
		// delete y, list is x->z
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
	// InsertBefore
	{
		list := newList(nil)
		a := list.RightPush("a")
		b := NewSingleListNode("b")
		// head
		assert.Equal(t, list.InsertAfter(a, b), true)
		equalList(t, list, equalData{
			Dump:    "a->b",
			HeadValue:    "a",
			HasHead: true,
			TailValue:    "b",
			HasTail: true,
			Length:  2,
		})
		assert.Equal(t, list.InsertAfter(NewSingleListNode("1"), NewSingleListNode("2")), false)
		// tail
		c := NewSingleListNode("c")
		assert.Equal(t, list.InsertAfter(b, c), true)
		equalList(t, list, equalData{
			Dump:    "a->b->c",
			HeadValue:    "a",
			HasHead: true,
			TailValue:    "c",
			HasTail: true,
			Length:  3,
		})
		// middle
		x := NewSingleListNode("x")
		assert.Equal(t, list.InsertAfter(b, x), true)
		equalList(t, list, equalData{
			Dump:    "a->b->x->c",
			HeadValue:    "a",
			HasHead: true,
			TailValue:    "c",
			HasTail: true,
			Length:  4,
		})
	}
	// FindByIndex
	{
		list := newList(nil)
		// empty
		{
			node, hasNode := list.FindByIndex(0)
			assert.Nil(t, node)
			assert.Equal(t,hasNode, false)
		}
		// one node
		a := list.RightPush("a")
		{
			node, hasNode := list.FindByIndex(0)
			assert.Equal(t, node, a)
			assert.Equal(t,hasNode, true)
		}
		// two node
		b := list.RightPush("b")
		{
			node, hasNode := list.FindByIndex(0)
			assert.Equal(t, node, a)
			assert.Equal(t,hasNode, true)
		}
		{
			node, hasNode := list.FindByIndex(1)
			assert.Equal(t, node, b)
			assert.Equal(t,hasNode, true)
		}
		// three node
		c := list.RightPush("c")
		{
			node, hasNode := list.FindByIndex(0)
			assert.Equal(t, node, a)
			assert.Equal(t,hasNode, true)
		}
		{
			node, hasNode := list.FindByIndex(1)
			assert.Equal(t, node, b)
			assert.Equal(t,hasNode, true)
		}
		{
			node, hasNode := list.FindByIndex(2)
			assert.Equal(t, node, c)
			assert.Equal(t,hasNode, true)
		}
		{
			node, hasNode := list.FindByIndex(3)
			assert.Nil(t, node)
			assert.Equal(t,hasNode, false)
		}
	}
	// Reverse
	{
		list := newList(nil)
		list.Reverse()
		assert.Equal(t, list.Dump(), "")
	}
	{
		list := newList(nil)
		list.RightPush("a")
		list.Reverse()
		assert.Equal(t, list.Dump(), "a")
	}
	{
		list := newList(nil)
		list.RightPush("a")
		list.RightPush("b")
		list.Reverse()
		assert.Equal(t, list.Dump(), "b->a")
	}
	{
		list := newList(nil)
		list.RightPush("a")
		list.RightPush("b")
		list.RightPush("c")
		list.Reverse()
		assert.Equal(t, list.Dump(), "c->b->a")
	}
	{
		list := newList(nil)
		list.RightPush("a")
		list.RightPush("b")
		list.RightPush("c")
		list.RightPush("d")
		list.Reverse()
		assert.Equal(t, list.Dump(), "d->c->b->a")
	}
	// cycle
	list := newList(nil)
	a := list.RightPush("a");_=a
	b := list.RightPush("b");_=b
	c := list.RightPush("c")
	d := list.RightPush("d");_=d
	assert.Equal(t, list.IsCycle(), false)
	e := list.RightPush("e")
	e.SetNext(c)
	assert.Equal(t, list.Dump(), "a->b->c->d->e->cycle{index:2, value:c}")
	assert.Equal(t, list.IsCycle(), true)
}

