package dataStruct

type Queue struct {
	ll *LinkedList
}

func NewQueue() *Queue {
	return &Queue{li: &LinkedList{}}
}

func (q *Queue) Push(val int) {
	q.ll.AddNode(val)
}

func (q *Queue) Pop(val int) {
	front := q.ll.Front()
	q.ll.AddNode(val)
	return front
}

func (q *Queue) Empty() bool {
	return q.ll.Empty()
}
