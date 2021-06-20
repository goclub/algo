package goclub_algo_stack

import (
	"strconv"
	"strings"
)

type ArrayStack struct {
	items []string
	len int
	cap int
}
func NewArrayStack(initialCap int) *ArrayStack {
	return &ArrayStack{
		cap: initialCap,
		items: make([]string, initialCap, initialCap),
	}
}
func (as *ArrayStack) Push(item string) {
	if as.len == as.cap {
		as.cap = as.cap*2
		newItems := make([]string, as.cap, as.cap)
		copy(newItems, as.items)
		as.items = newItems
	}
	// go array append 会自动扩容,为了演示算法不使用 append
	as.items[as.len] = item
	as.len++
}
func (as *ArrayStack) Pop() (s string, ok bool) {
	if as.len == 0 {
		return "", false
	}
	as.len--
	s = as.items[as.len]
	if as.len < as.cap / 2 {
		as.cap = as.cap/2
		newItems := make([]string, as.cap, as.cap)
		copy(newItems, as.items)
		as.items = newItems
	}
	return s,true
}
func (as ArrayStack) Len() int {
	return as.len
}
func (as ArrayStack) Cap() int {
	return as.cap
}
func (as ArrayStack) Dump () (s string) {
	// 测试用
	if as.cap != cap(as.items) {
		panic("cap error")
	}
	return "ArrayStack:" +
		" cap(" + strconv.FormatUint(uint64(as.cap), 10) + ")" +
		" len(" +strconv.FormatUint(uint64(as.len), 10)+")" +
		" items("+strings.Join(as.items, ",")+")"
}