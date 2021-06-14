package nimoc_algo_geeekbang_linkedlist


type SingleList2 struct {

}

func (s SingleList2) Head() (head *SingleListNode, hasHead bool) {
	panic("implement me")
}

func (s SingleList2) Tail() (tail *SingleListNode, hasTail bool) {
	panic("implement me")
}

func (s SingleList2) Length() int {
	panic("implement me")
}

func (s SingleList2) Dump() string {
	panic("implement me")
}

func (s SingleList2) LeftPush(v interface{}) (newNode *SingleListNode) {
	panic("implement me")
}

func (s SingleList2) RightPush(v interface{}) (newNode *SingleListNode) {
	panic("implement me")
}

func (s SingleList2) PrevNode(pivot *SingleListNode) (prevNode *SingleListNode, hasPrevNode bool) {
	panic("implement me")
}

func (s SingleList2) DeleteByNode(node *SingleListNode) (ok bool) {
	panic("implement me")
}

func (s SingleList2) InsertBefore(pivot *SingleListNode, newNode *SingleListNode) (ok bool) {
	panic("implement me")
}

func (s SingleList2) InsertAfter(pivot *SingleListNode, newNode *SingleListNode) (ok bool) {
	panic("implement me")
}

func (s SingleList2) FindByIndex(index int) (target *SingleListNode, hasNode bool) {
	panic("implement me")
}

func (s SingleList2) Reverse() {
	panic("implement me")
}

func (s SingleList2) IsCycle() bool {
	panic("implement me")
}

func NewSingleList2(node *SingleListNode) SingleLister {
	return SingleList2{}
}
