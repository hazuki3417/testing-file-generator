package openapi

type ErrorPositionQueue struct {
	data []ErrorPosition
}

func NewErrorPositionQueue(cap int) *ErrorPositionQueue {
	return &ErrorPositionQueue{data: make([]ErrorPosition, 0, cap)}
}

func (q *ErrorPositionQueue) Enqueue(data ErrorPosition) {
	q.data = append(q.data, data)
}

func (q *ErrorPositionQueue) Denqueue() ErrorPosition {
	deferFunc := func() {
		q.data = q.data[1:]
	}
	defer deferFunc()

	return q.data[0]
}

func (q *ErrorPositionQueue) All() []ErrorPosition {
	deferFunc := func() {
		q.data = make([]ErrorPosition, 0, len(q.data))
	}
	defer deferFunc()

	return q.data[:]
}

func (q *ErrorPositionQueue) Size() int {
	return len(q.data)
}

func (q *ErrorPositionQueue) IsEmpty() bool {
	return q.Size() == 0
}

func (q *ErrorPositionQueue) IsNotEmpty() bool {
	return 0 < q.Size()
}

func (q *ErrorPositionQueue) HasEnqueue() bool {
	return q.IsEmpty()
}

func (q *ErrorPositionQueue) HasDenqueue() bool {
	return q.IsNotEmpty()
}
