package queue

type stringQueue struct {
	data []string
}

func NewTypeString(cap int) *stringQueue {
	return &stringQueue{data: make([]string, 0, cap)}
}

func (q *stringQueue) Enqueue(data string) {
	q.data = append(q.data, data)
}

func (q *stringQueue) Denqueue() string {
	deferFunc := func() {
		q.data = q.data[1:]
	}
	defer deferFunc()

	return q.data[0]
}

func (q *stringQueue) All() []string {
	deferFunc := func() {
		q.data = make([]string, 0, len(q.data))
	}
	defer deferFunc()

	return q.data[:]
}

func (q *stringQueue) Size() int {
	return len(q.data)
}

func (q *stringQueue) IsEmpty() bool {
	return q.Size() == 0
}

func (q *stringQueue) IsNotEmpty() bool {
	return 0 < q.Size()
}

func (q *stringQueue) HasEnqueue() bool {
	return q.IsEmpty()
}

func (q *stringQueue) HasDenqueue() bool {
	return q.IsNotEmpty()
}
