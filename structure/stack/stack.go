package stack

type StackEntry struct {
	element []interface{}
}

func InitStackEntry() *StackEntry {
	return &StackEntry{}
}

//进栈
func (se *StackEntry) Push(e interface{}) {
	se.element = append(se.element, e)
}

//出栈
func (se *StackEntry) Pop() interface{} {
	size := se.Size()
	if size == 0 {
		return nil
	}
	lastElement := se.element[size-1]
	se.element[size-1] = nil
	se.element = se.element[:size-1]
	return lastElement
}

//栈顶元素
func (se *StackEntry) Top() interface{} {
	size := se.Size()
	if size == 0 {
		return nil
	}
	return se.element[size-1]
}

//栈长度
func (se *StackEntry) Size() int {
	return len(se.element)
}

//栈是否为空
func (se *StackEntry) IsEmpty() bool {
	return len(se.element) == 0
}

//全部清除栈内数据
func (se *StackEntry) Clear() bool {
	if se.IsEmpty() {
		return false
	}
	for i := 0; i < se.Size(); i++ {
		se.element[i] = nil
	}
	se.element = make([]interface{}, 0)
	return true
}
