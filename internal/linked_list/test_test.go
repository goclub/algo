package goclub_algo_linked_list

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

type equalData struct {
	FirstValue interface{}
	HasFirst bool
	LastValue interface{}
	HasLast bool
	Size int
	Dump string
}
func equalList(t *testing.T, list SingleLister, data equalData) {
	// Dump
	assert.Equal(t, data.Dump, list.Dump())
	// first
	{
		first, hasFirst := list.First()
		assert.Equal(t, data.HasFirst, hasFirst)
		if hasFirst {
			assert.Equal(t, data.FirstValue, first.Value())
		} else {
			assert.Equal(t, data.FirstValue, nil)
		}
	}
	// last
	{
		last, hasLast := list.Last()
		assert.Equal(t, data.HasLast, hasLast)
		if hasLast {
			assert.Equal(t, data.LastValue, last.Value())
		} else {
			assert.Equal(t, data.LastValue, nil)
		}
	}
	// length
	assert.Equal(t, data.Size, list.Size())
}



type TestSingleListSuite struct {
	suite.Suite
	NewList func(node *SingleListNode) SingleLister
}

func (suite TestSingleListSuite) TestEmptyList() {
	t := suite.T()
	newList := suite.NewList
	{
		list := newList(nil)
		equalList(t, list, equalData{
			FirstValue:    nil,
			HasFirst: false,
			LastValue:    nil,
			HasLast: false,
			Size:  0,
			Dump:    "",
		})
	}
}

func (suite TestSingleListSuite) TestAddFirst() {
	t := suite.T()
	newList := suite.NewList
	{
		list := newList(NewSingleListNode("a"))
		equalList(t, list, equalData{
			FirstValue:    "a",
			HasFirst: true,
			LastValue:    "a",
			HasLast: true,
			Size:  1,
			Dump:    "a",
		})
		list.AddFirst("b")
		equalList(t, list, equalData{
			Dump:    "b->a",
			FirstValue:    "b",
			HasFirst: true,
			LastValue:    "a",
			HasLast: true,
			Size:  2,
		})
		list.AddFirst("c")
		equalList(t, list, equalData{
			Dump:    "c->b->a",
			FirstValue:    "c",
			HasFirst: true,
			LastValue:    "a",
			HasLast: true,
			Size:  3,
		})
	}
}
func (suite TestSingleListSuite) TestAddLast() {
	t := suite.T()
	newList := suite.NewList
	list := newList(NewSingleListNode("a"))
	equalList(t, list, equalData{
		FirstValue:    "a",
		HasFirst: true,
		LastValue:    "a",
		HasLast: true,
		Size:  1,
		Dump:    "a",
	})
	list.AddLast("b")
	equalList(t, list, equalData{
		Dump:    "a->b",
		FirstValue:    "a",
		HasFirst: true,
		LastValue:    "b",
		HasLast: true,
		Size:  2,
	})
	list.AddLast("c")
	equalList(t, list, equalData{
		Dump:    "a->b->c",
		FirstValue:    "a",
		HasFirst: true,
		LastValue:    "c",
		HasLast: true,
		Size:  3,
	})
}
func (suite TestSingleListSuite) TestPrevNode() {
	t := suite.T()
	newList := suite.NewList
	list := newList(nil)
	a := list.AddLast("a")
	b := list.AddLast("b")
	c := list.AddLast("c")
	equalList(t, list, equalData{
		Dump:    "a->b->c",
		FirstValue:    "a",
		HasFirst: true,
		LastValue:    "c",
		HasLast: true,
		Size:  3,
	})
	// first
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
	// last
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

func (suite TestSingleListSuite) TestRemoveByNode() {
	t := suite.T()
	newList := suite.NewList
	list := newList(nil)
	a := list.AddLast("a")
	b := list.AddLast("b")
	c := list.AddLast("c")
	equalList(t, list, equalData{
		Dump:    "a->b->c",
		FirstValue:    "a",
		HasFirst: true,
		LastValue:    "c",
		HasLast: true,
		Size:  3,
	})
	// first
	{
		assert.Equal(t, list.RemoveByNode(a), true)
		equalList(t, list, equalData{
			Dump:    "b->c",
			FirstValue:    "b",
			HasFirst: true,
			LastValue:    "c",
			HasLast: true,
			Size:  2,
		})
		assert.Equal(t, list.RemoveByNode(a), false)
	}
	// last
	{
		assert.Equal(t, list.RemoveByNode(c), true)
		equalList(t, list, equalData{
			Dump:    "b",
			FirstValue:    "b",
			HasFirst: true,
			LastValue:    "b",
			HasLast: true,
			Size:  1,
		})
	}
	// one node
	{
		assert.Equal(t, list.RemoveByNode(b), true)
		equalList(t, list, equalData{
			Dump:    "",
			FirstValue:    nil,
			HasFirst: false,
			LastValue:    nil,
			HasLast: false,
			Size:  0,
		})
	}
	// middle
	// x->y->z
	// delete y, list is x->z
	{
		_ = list.AddLast("x")
		y := list.AddLast("y")
		_ = list.AddLast("z")
		equalList(t, list, equalData{
			Dump:    "x->y->z",
			FirstValue:    "x",
			HasFirst: true,
			LastValue:    "z",
			HasLast: true,
			Size:  3,
		})
		assert.Equal(t, list.RemoveByNode(y), true)
		equalList(t, list, equalData{
			Dump:    "x->z",
			FirstValue:    "x",
			HasFirst: true,
			LastValue:    "z",
			HasLast: true,
			Size:  2,
		})
	}
}

func (suite TestSingleListSuite) TestInsertBefore() {
	t := suite.T()
	newList := suite.NewList
	list := newList(nil)
	a := list.AddLast("a")
	// first
	b := NewSingleListNode("b")
	assert.Equal(t, list.InsertBefore(a, b), true)
	equalList(t, list, equalData{
		Dump:    "b->a",
		FirstValue:    "b",
		HasFirst: true,
		LastValue:    "a",
		HasLast: true,
		Size:  2,
	})
	c := NewSingleListNode("c")
	assert.Equal(t, list.InsertBefore(b, c), true)
	equalList(t, list, equalData{
		Dump:    "c->b->a",
		FirstValue:    "c",
		HasFirst: true,
		LastValue:    "a",
		HasLast: true,
		Size:  3,
	})
	// middle
	x := NewSingleListNode("x")
	assert.Equal(t, list.InsertBefore(b, x), true)
	equalList(t, list, equalData{
		Dump:    "c->x->b->a",
		FirstValue:    "c",
		HasFirst: true,
		LastValue:    "a",
		HasLast: true,
		Size:  4,
	})
	// last
	z := NewSingleListNode("z")
	assert.Equal(t, list.InsertBefore(a, z), true)
	equalList(t, list, equalData{
		Dump:    "c->x->b->z->a",
		FirstValue:    "c",
		HasFirst: true,
		LastValue:    "a",
		HasLast: true,
		Size:  5,
	})
}
func (suite TestSingleListSuite) TestInsertAfter() {
	t := suite.T()
	newList := suite.NewList
	list := newList(nil)
	a := list.AddLast("a")
	b := NewSingleListNode("b")
	// first
	assert.Equal(t, list.InsertAfter(a, b), true)
	equalList(t, list, equalData{
		Dump:    "a->b",
		FirstValue:    "a",
		HasFirst: true,
		LastValue:    "b",
		HasLast: true,
		Size:  2,
	})
	assert.Equal(t, list.InsertAfter(NewSingleListNode("1"), NewSingleListNode("2")), false)
	// last
	c := NewSingleListNode("c")
	assert.Equal(t, list.InsertAfter(b, c), true)
	equalList(t, list, equalData{
		Dump:    "a->b->c",
		FirstValue:    "a",
		HasFirst: true,
		LastValue:    "c",
		HasLast: true,
		Size:  3,
	})
	// middle
	x := NewSingleListNode("x")
	assert.Equal(t, list.InsertAfter(b, x), true)
	equalList(t, list, equalData{
		Dump:    "a->b->x->c",
		FirstValue:    "a",
		HasFirst: true,
		LastValue:    "c",
		HasLast: true,
		Size:  4,
	})
}
func (suite TestSingleListSuite) TestFindByIndex() {
	t := suite.T()
	newList := suite.NewList
	list := newList(nil)
	// empty
	{
		node, hasNode := list.FindByIndex(0)
		assert.Nil(t, node)
		assert.Equal(t,hasNode, false)
	}
	// one node
	a := list.AddLast("a")
	{
		node, hasNode := list.FindByIndex(0)
		assert.Equal(t, node, a)
		assert.Equal(t,hasNode, true)
	}
	// two node
	b := list.AddLast("b")
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
	c := list.AddLast("c")
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
func (suite TestSingleListSuite) TestReverse() {
	t := suite.T()
	newList := suite.NewList
	{
		list := newList(nil)
		list.Reverse()
		assert.Equal(t, list.Dump(), "")
	}
	{
		list := newList(nil)
		list.AddLast("a")
		list.Reverse()
		assert.Equal(t, list.Dump(), "a")
	}
	{
		list := newList(nil)
		list.AddLast("a")
		list.AddLast("b")
		list.Reverse()
		assert.Equal(t, list.Dump(), "b->a")
	}
	{
		list := newList(nil)
		list.AddLast("a")
		list.AddLast("b")
		list.AddLast("c")
		list.Reverse()
		assert.Equal(t, list.Dump(), "c->b->a")
	}
	{
		list := newList(nil)
		list.AddLast("a")
		list.AddLast("b")
		list.AddLast("c")
		list.AddLast("d")
		list.Reverse()
		assert.Equal(t, list.Dump(), "d->c->b->a")
	}
}
func (suite TestSingleListSuite) TestIsCycle() {
	t := suite.T()
	newList := suite.NewList
	list := newList(nil)
	a := list.AddLast("a");_=a
	b := list.AddLast("b");_=b
	c := list.AddLast("c")
	d := list.AddLast("d");_=d
	assert.Equal(t, list.IsCycle(), false)
	e := list.AddLast("e")
	e.SetNext(c)
	assert.Equal(t, list.Dump(), "a->b->c->d->e->cycle{index:2, value:c}")
	assert.Equal(t, list.IsCycle(), true)
}

