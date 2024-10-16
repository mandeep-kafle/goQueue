package queue

type Queue struct {
	arr []interface{}
}

func (q *Queue) Len() interface{} { return len(q.arr) }

func (q *Queue) Push(v interface{}) {
	q.arr = append(q.arr, v)
}

func (q *Queue) Pop() {
	q.arr = q.arr[1:]

}
func (q *Queue) Front() interface{} {
	return q.arr[0]
}
func NewQueue() *Queue { return &Queue{} }
