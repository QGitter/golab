package linelist

import "errors"

//定义单循环链表的节点数据结构
type CircleNode struct {
	data interface{}
	next *CircleNode
}

//定义单循环链表的数据结构
type CircleList struct {
	tail *CircleNode
	size int
}

func InitCircleList() *CircleList {
	return &CircleList{tail: nil, size: 0}
}

func InitCircleNode(data interface{}) *CircleNode {
	return &CircleNode{data: data, next: nil}
}

//单链表在表尾添加数据
func (cl *CircleList) Append(data *CircleNode) bool {
	if data == nil {
		return false
	}
	if cl.size == 0 {
		data.next = data
	} else {
		curNode := cl.tail.next
		data.next = curNode
		cl.tail.next = data
	}
	cl.tail = data
	cl.size++
	return true
}

//单循环链表插入数据
func (cl *CircleList) Insert(num int, data *CircleNode) error {
	if data == nil {
		return errors.New("要插入的节点数据为空")
	}
	if cl.size == 0 || cl.size == num {
		cl.Append(data)
	} else {
		var curNode *CircleNode
		if num == 0 {
			curNode = cl.tail
		} else {
			curNode = cl.Get(num)
			if cl.size == num {
				cl.tail = data
			}
		}
		data.next = curNode.next
		curNode.next = data
		cl.size++
	}
	return nil
}

//单循环链表查询数据
func (cl *CircleList) Get(num int) *CircleNode {
	if num < 0 || num > cl.size-1 {
		return nil
	}
	curNode := cl.tail
	for i := 0; i < num; i++ {
		curNode = curNode.next
	}
	return curNode
}

//单循环链表查询全部数据
func (cl *CircleList) GetAll() []interface{} {
	result := make([]interface{}, 0)
	curNode := cl.tail
	for i := 0; i < cl.size; i++ {
		result = append(result, curNode.data)
		curNode = curNode.next
	}
	return result
}

//单循环链表按序号删除数据
func (cl *CircleList) RemoveInt(num int) error {
	if cl.size == 0 {
		return errors.New("循环链表为空")
	}
	if num > cl.size-1 {
		return errors.New("越界")
	}

	if cl.size == 1 {
		cl.tail = nil
		cl.size = 0
		return nil
	} else {
		var curNode *CircleNode
		var data *CircleNode
		if num == 0 {
			curNode = cl.tail
		} else {
			curNode = cl.Get(num - 1)
		}

		data = curNode.next
		curNode.next = data.next

		if num == cl.size-1 {
			cl.tail = curNode
		}

		data.next = nil
		data = nil
		cl.size--

		return nil
	}
}

//单循环链表删除全部数据
func (cl *CircleList) RemoveAll() bool {
	if cl.size == 0 {
		return false
	}

	for i := 0; i < cl.size; i++ {
		curNode := cl.tail
		cl.tail = curNode.next
		curNode.next = nil
	}
	cl.tail = nil
	cl.size = 0

	return true
}
