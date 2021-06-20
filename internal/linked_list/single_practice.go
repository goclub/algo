package goclub_algo_linked_list


// 单链表练习模板
type SingleListPractice struct {

}


func NewSingleListPractice(node *SingleListNode) SingleLister {
	return SingleListPractice{}
}

func (s SingleListPractice) First() (first *SingleListNode, hasFirst bool) {
	panic("implement me")
}

func (s SingleListPractice) Last() (last *SingleListNode, hasLast bool) {
	panic("implement me")
}

func (s SingleListPractice) Size() int {
	panic("implement me")
}

func (s SingleListPractice) Dump() string {
	panic("implement me")
}

func (s SingleListPractice) AddFirst(v interface{}) (newNode *SingleListNode) {
	panic("implement me")
}

func (s SingleListPractice) AddLast(v interface{}) (newNode *SingleListNode) {
	panic("implement me")
}

func (s SingleListPractice) PrevNode(pivot *SingleListNode) (prevNode *SingleListNode, hasPrevNode bool) {
	panic("implement me")
}

func (s SingleListPractice) RemoveByNode(node *SingleListNode) (ok bool) {
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
