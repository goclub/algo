package goclub_algo_linked_list


// 单链表练习模板
type SingleListPractice struct {

}


func NewSingleListPractice(node *SingleListNode) SingleLister {
	return SingleListPractice{}
}

func (s SingleListPractice) Head() (head *SingleListNode, hasHead bool) {
	panic("implement me")
}

func (s SingleListPractice) Tail() (tail *SingleListNode, hasTail bool) {
	panic("implement me")
}

func (s SingleListPractice) Length() int {
	panic("implement me")
}

func (s SingleListPractice) Dump() string {
	panic("implement me")
}

func (s SingleListPractice) LeftPush(v interface{}) (newNode *SingleListNode) {
	panic("implement me")
}

func (s SingleListPractice) RightPush(v interface{}) (newNode *SingleListNode) {
	panic("implement me")
}

func (s SingleListPractice) PrevNode(pivot *SingleListNode) (prevNode *SingleListNode, hasPrevNode bool) {
	panic("implement me")
}

func (s SingleListPractice) DeleteByNode(node *SingleListNode) (ok bool) {
	panic("implement me")
}

func (s SingleListPractice) InsertBefore(pivot *SingleListNode, newNode *SingleListNode) (ok bool) {
	panic("implement me")
}

func (s SingleListPractice) InsertAfter(pivot *SingleListNode, newNode *SingleListNode) (ok bool) {
	panic("implement me")
}

func (s SingleListPractice) FindByIndex(index int) (target *SingleListNode, hasNode bool) {
	panic("implement me")
}

func (s SingleListPractice) Reverse() {
	panic("implement me")
}

func (s SingleListPractice) IsCycle() bool {
	panic("implement me")
}
