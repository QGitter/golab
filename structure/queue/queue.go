package queue

type QueueEntry struct {
	element []interface{}
}

func InitQueueEntry() *QueueEntry {
	return &QueueEntry{}
}

func (qe *QueueEntry) Len() int {
	return len(qe.element)
}

func (qe *QueueEntry) Push(data interface{}) {
	qe.element = append(qe.element, data)
}

func (qe *QueueEntry) Pop() interface{} {
	if len(qe.element) == 0 {
		return nil
	}
	data := qe.element[0]
	qe.element[0] = nil
	qe.element = qe.element[1:len(qe.element)]
	return data
}

func (qe *QueueEntry) Top() interface{} {
	if len(qe.element) == 0 {
		return nil
	}
	return qe.element[0]
}

func (qe *QueueEntry) IsEmpty() bool {
	return len(qe.element) == 0
}
