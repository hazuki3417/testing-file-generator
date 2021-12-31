package openapi

type ErrorQueue struct {
	data []Error
}

func NewErrorQueue(cap int) *ErrorQueue {
	return &ErrorQueue{data: make([]Error, 0, cap)}
}

func (q *ErrorQueue) Enqueue(data Error) {
	q.data = append(q.data, data)
}

func (q *ErrorQueue) Denqueue() Error {
	deferFunc := func() {
		q.data = q.data[1:]
	}
	defer deferFunc()

	return q.data[0]
}

func (q *ErrorQueue) All() []Error {
	deferFunc := func() {
		q.data = make([]Error, 0, len(q.data))
	}
	defer deferFunc()

	return q.data[:]
}

func (q *ErrorQueue) Size() int {
	return len(q.data)
}

func (q *ErrorQueue) IsEmpty() bool {
	return q.Size() == 0
}

func (q *ErrorQueue) IsNotEmpty() bool {
	return 0 < q.Size()
}

func (q *ErrorQueue) HasEnqueue() bool {
	return q.IsEmpty()
}

func (q *ErrorQueue) HasDenqueue() bool {
	return q.IsNotEmpty()
}
