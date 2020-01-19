package linelist

// 单链表结点
type SingleList struct {
	Data interface{} //单链表的数据域
	Next *SingleList //单链表的指针域
}

func NewSingleList() *SingleList {
	return &SingleList{Data: "", Next: nil}
}

type SingleListr interface {
	GetFirst() *SingleList
	GetLast() *SingleList
	Length() int
	Add(data interface{}) bool
	GetElem(index int) (interface{}, error)
	Delete(index int) bool
}

//返回第一个结点
func (this *SingleList) GetFirst() *SingleList {
	if this.Next == nil {
		return nil
	}
	return this.Next
}

//返回最后一个结点
func (this *SingleList) GetLast() *SingleList {
	if this.Next == nil {
		return nil
	}
	point := this
	for point.Next != nil {
		point = point.Next
	}
	if point.Next == nil {
		return point
	}
	return nil
}

//获取单链表的长度
func (this *SingleList) Length() int {
	point := this
	length := 0

	for point.Next != nil {
		length++
		point = point.Next
	}
	return length
}

//往单链表的末尾加一个元素
func (this *SingleList) Add(data interface{}) bool {
	point := this
	for point.Next != nil {
		point = point.Next
	}
	tmpSingle := SingleList{Data: data}
	point.Next = &tmpSingle
	return true
}

//获取所有结点的值
func (this *SingleList) GetAll() []interface{} {
	result := make([]interface{}, 0)
	point := this
	for point.Next != nil {
		result = append(result, point.Data)
		point = point.Next
	}
	result = append(result, point.Data)
	return result
}

//获取索引为index的结点
func (this *SingleList) GetElem(index int) *SingleList {
	point := this
	if index < 0 || index > this.Length() {
		panic("check index error")
		return nil
	}
	for i := 0; i < index; i++ {
		point = point.Next
	}
	return point
}

//删除第index个结点
func (this *SingleList) Delete(index int) bool {
	if index < 0 || index > this.Length() {
		panic("please check index")
		return false
	}
	point := this
	for i := 0; i < index-1; i++ {
		point = point.Next
	}
	point.Next = point.Next.Next
	return true
}
