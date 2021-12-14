package queue

type intQueue struct {
	data []int
}

func NewTypeInt(cap int) *intQueue {
	return &intQueue{data: make([]int, 0, cap)}
}

func (q *intQueue) Enqueue(data int) {
	q.data = append(q.data, data)
}

func (q *intQueue) Denqueue() int {
	deferFunc := func() {
		q.data = q.data[1:]
	}
	defer deferFunc()

	return q.data[0]
}

func (q *intQueue) All() []int {
	deferFunc := func() {
		q.data = make([]int, 0, len(q.data))
	}
	defer deferFunc()

	return q.data[:]
}
func (q *intQueue) Size() int {
	return len(q.data)
}

func (q *intQueue) IsEmpty() bool {
	return q.Size() == 0
}

func (q *intQueue) IsNotEmpty() bool {
	return 0 < q.Size()
}

func (q *intQueue) HasEnqueue() bool {
	return q.IsEmpty()
}

func (q *intQueue) HasDenqueue() bool {
	return q.IsNotEmpty()
}
