package nimoc_algo_geeekbang_linkedlist

func NewSingleListNode (v interface{}) *SingleListNode {
	return &SingleListNode{
		value: v,
		next: nil,
	}
}
type SingleListNode struct {
	value interface{}
	next *SingleListNode
}
func (node *SingleListNode) Next() (nextNode *SingleListNode, hasNext bool) {
	return node.next,  node.next != nil
}
func (node *SingleListNode) Value() interface{} {
	return node.value
}

type SingleLinkedLister interface {
	// 获取头节点
	Head() (head *SingleListNode, hasHead bool)
	// 获取尾节点
	Tail() (tail *SingleListNode, hasTail bool)
	// 获取链表长度
	Length () int
	// 打印链表
	Dump() string
	// // 头入
	// LeftPush(v interface{})
	// // 尾入
	// RightPush(v interface{})
	// // 查找节点的上一个节点 很多方法都是基于 PrevNode 实现的所以需要优先实现 PrevNode
	// PrevNode(pivot *SingleListNode) (prevNode *SingleListNode, hasPrevNode bool)
	// // 根据节点删除节点
	// DeleteByNode(node *SingleListNode) (ok bool)
	// // 在指定节点之前插入 a->b.InsertAfter(a,z) = z->a->b
	// InsertBefore(pivot *SingleListNode, element interface{}) (ok bool)
	// // 在指定节点之后插入 a->b.InsertAfter(a,z) = a->z->b
	// InsertAfter(pivot *SingleListNode, element interface{}) (ok bool)
	// // 根据 index 查询节点
	// FindByIndex(index uint) (node *SingleListNode, hasNode bool)
}
