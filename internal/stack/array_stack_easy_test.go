package goclub_algo_stack

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewArrayStack(t *testing.T) {
	a := NewArrayStack(2)
	assert.Equal(t, a.Dump(), "ArrayStack: cap(2) len(0) items(,)")
	a.Push("a")
	assert.Equal(t, a.Dump(), "ArrayStack: cap(2) len(1) items(a,)")
	a.Push("b")
	assert.Equal(t, a.Dump(), "ArrayStack: cap(2) len(2) items(a,b)")
	a.Push("c")
	assert.Equal(t, a.Dump(), "ArrayStack: cap(4) len(3) items(a,b,c,)")
	a.Push("d")
	assert.Equal(t, a.Dump(), "ArrayStack: cap(4) len(4) items(a,b,c,d)")
	a.Push("e")
	assert.Equal(t, a.Dump(), "ArrayStack: cap(8) len(5) items(a,b,c,d,e,,,)")
	a.Push("f")
	a.Push("g")
	a.Push("h")
	a.Push("i")
	a.Push("j")
	a.Push("k")
	assert.Equal(t, a.Dump(), "ArrayStack: cap(16) len(11) items(a,b,c,d,e,f,g,h,i,j,k,,,,,)")
	equalPop(t, a, "k", true)
	equalPop(t, a, "j", true)
	equalPop(t, a, "i", true)
	assert.Equal(t, a.Dump(), "ArrayStack: cap(16) len(8) items(a,b,c,d,e,f,g,h,i,j,k,,,,,)")
	equalPop(t, a, "h", true)
	// 因为自动缩容会变成 cap(8)
	assert.Equal(t, a.Dump(), "ArrayStack: cap(8) len(7) items(a,b,c,d,e,f,g,h)")
	equalPop(t, a, "g", true)
	// pop时无需删除数据 所以 gh依然存在
	assert.Equal(t, a.Dump(), "ArrayStack: cap(8) len(6) items(a,b,c,d,e,f,g,h)")
	equalPop(t, a, "f", true)
	equalPop(t, a, "e", true)
	assert.Equal(t, a.Dump(), "ArrayStack: cap(8) len(4) items(a,b,c,d,e,f,g,h)")
	equalPop(t, a, "d", true)
	assert.Equal(t, a.Dump(), "ArrayStack: cap(4) len(3) items(a,b,c,d)")
	equalPop(t, a, "c", true)
	equalPop(t, a, "b", true)
	equalPop(t, a, "a", true)
	assert.Equal(t, a.Dump(), "ArrayStack: cap(1) len(0) items(a)")
	equalPop(t, a, "", false)
}

func equalPop(t *testing.T, as *ArrayStack, s string, ok bool) {
	s2, ok2:=as.Pop()
	assert.Equal(t, s2, s)
	assert.Equal(t, ok2, ok)
}