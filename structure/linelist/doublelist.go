package linelist

import (
	"errors"
)

var (
	NUMERROR = errors.New("链表越界")
)
//定义双向链表节点结构体
type DoubleNode struct {
	data interface{}
	prev *DoubleNode
	next *DoubleNode
}

//定义双向链表结构体
type DoubleList struct {
	head *DoubleNode
	tail *DoubleNode
	size int
}

//初始化链表

func InitDoubleList() *DoubleList {
	return &DoubleList{head: nil, tail: nil, size: 0}
}

func InitDoubleNode(data interface{}) *DoubleNode {
	return &DoubleNode{data: data, prev: nil, next: nil}
}

//获取链表的长度
func (dl *DoubleList) GetSize() int {
	return dl.size
}

//获取链表头部节点
func (dl *DoubleList) GetHead() *DoubleNode {
	return dl.head
}

//获取链表尾部节点
func (dl *DoubleList) GetTail() *DoubleNode {
	return dl.tail
}

//在头部追加节点
func (dl *DoubleList) AddHeadNode(node *DoubleNode) int {
	if dl.GetSize() == 0 {
		dl.head = node
		dl.tail = node
		node.prev = nil
		node.next = nil
	} else {
		dl.head.prev = node
		node.prev = nil
		node.next = dl.head
		dl.head = node
	}
	dl.size += 1
	return dl.size
}

//在尾部追加节点
func (dl *DoubleList) AddTailNode(node *DoubleNode) int {
	if dl.GetSize() == 0 {
		dl.head = node
		dl.tail = node
		node.prev = nil
		node.next = nil
	} else {
		dl.tail.next = node
		node.prev = dl.tail
		node.next = nil
		dl.tail = node
	}
	dl.size += 1
	return dl.size
}

//在链表某个序号之后插入节点
func (dl *DoubleList) InsertNextInt(num int, data *DoubleNode) bool {
	if data == nil || num > dl.GetSize()-1 || num < 0 {
		return false
	}
	switch {
	case dl.GetSize() == 0:
		dl.AddHeadNode(data)
	case num == dl.GetSize()-1:
		dl.AddTailNode(data)
	default:
		curNode, err := dl.GetOrder(num)
		if err != nil {
			return false
		}
		data.prev = curNode
		data.next = curNode.next
		curNode.next = data
		curNode.next.prev = data
		dl.size++
	}
	return true
}

//顺序查询某个序号的数据
func (dl *DoubleList) GetOrder(num int) (*DoubleNode, error) {
	switch {
	case dl.GetSize() == 0:
		return nil, NUMERROR
	case num == 0:
		return dl.head, nil
	case num > dl.GetSize()-1:
		return nil, NUMERROR
	case num == dl.GetSize()-1:
		return dl.tail, nil
	default:
		data := dl.head
		for i := 0; i < num; i++ {
			data = data.next
		}
		return data, nil
	}
}

//倒序查询某个序号数据
func (dl *DoubleList) GetReverse(num int) (data *DoubleNode, err error) {
	switch {
	case num == 0:
		data = dl.tail
	case num > dl.GetSize()-1:
		err = NUMERROR
	case num == dl.GetSize()-1:
		data = dl.head
	default:
		data = dl.tail
		for i := 0; i < num; i++ {
			data = data.prev
		}
	}
	return
}

//获取链表中所有数据
func (dl *DoubleList) GetAll() []interface{} {
	result := make([]interface{}, 0)
	if dl.GetSize() == 0 {
		return nil
	}
	curNode := dl.head
	for i := 0; i < dl.GetSize(); i++ {
		result = append(result, curNode.data)
		curNode = curNode.next
	}
	return result
}

//删除某个序号的数据
func (dl *DoubleList) Remove(num int) error {
	if dl.GetSize() == 0 {
		return NUMERROR
	}

	var curNode *DoubleNode
	var err error
	if curNode, err = dl.GetOrder(num); err != nil {
		return err
	}

	if num == 0 {
		curNode.next.prev = nil
		dl.head = curNode.next
	} else if num == dl.size-1 {
		curNode.prev.next = nil
		dl.tail = curNode.prev
	} else {
		curNode.prev.next = curNode.next
		curNode.next.prev = curNode.prev
	}

	curNode.prev = nil
	curNode.next = nil
	dl.size--
	return nil
}

//删除链表中的全部数据
func (dl *DoubleList) RemoveAll() bool {
	for i := 0; i < dl.GetSize(); i++ {
		curNode := dl.head
		dl.head = curNode.next
		curNode.next = nil
		curNode.prev = nil
	}
	dl.tail = nil
	dl.size = 0
	return true
}
