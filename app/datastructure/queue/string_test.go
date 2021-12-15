package queue_test

import (
	"testing"

	queue "github.com/hazuki3417/testing-file-generator/datastructure/queue"
	"github.com/stretchr/testify/assert"
)

func Test_StringQueue(t *testing.T) {
	queue := queue.NewTypeString(5)

	assert.EqualValues(t, queue.IsEmpty(), true)
	assert.EqualValues(t, queue.IsNotEmpty(), false)
	assert.EqualValues(t, queue.HasEnqueue(), true)
	assert.EqualValues(t, queue.HasDenqueue(), false)
	assert.EqualValues(t, queue.Size(), 0)

	data1 := "string1"
	queue.Enqueue(data1)

	assert.EqualValues(t, queue.IsEmpty(), false)
	assert.EqualValues(t, queue.IsNotEmpty(), true)
	assert.EqualValues(t, queue.HasEnqueue(), false)
	assert.EqualValues(t, queue.HasDenqueue(), true)
	assert.EqualValues(t, queue.Size(), 1)

	data2 := "string2"
	queue.Enqueue(data2)

	assert.EqualValues(t, queue.IsEmpty(), false)
	assert.EqualValues(t, queue.IsNotEmpty(), true)
	assert.EqualValues(t, queue.HasEnqueue(), false)
	assert.EqualValues(t, queue.HasDenqueue(), true)
	assert.EqualValues(t, queue.Size(), 2)

	assert.EqualValues(t, data1, queue.Denqueue())
	assert.EqualValues(t, queue.IsEmpty(), false)
	assert.EqualValues(t, queue.IsNotEmpty(), true)
	assert.EqualValues(t, queue.HasEnqueue(), false)
	assert.EqualValues(t, queue.HasDenqueue(), true)
	assert.EqualValues(t, queue.Size(), 1)

	assert.EqualValues(t, data2, queue.Denqueue())
	assert.EqualValues(t, queue.IsEmpty(), true)
	assert.EqualValues(t, queue.IsNotEmpty(), false)
	assert.EqualValues(t, queue.HasEnqueue(), true)
	assert.EqualValues(t, queue.HasDenqueue(), false)
	assert.EqualValues(t, queue.Size(), 0)

	data3 := "string3"
	queue.Enqueue(data3)
	data4 := "string4"
	queue.Enqueue(data4)
	data5 := "string5"
	queue.Enqueue(data5)

	assert.EqualValues(t, queue.Size(), 3)
	assert.EqualValues(t, queue.All(), []string{data3, data4, data5})
	assert.EqualValues(t, queue.Size(), 0)

	assert.EqualValues(t, queue.IsExists(data1), false)
	queue.Enqueue(data1)
	assert.EqualValues(t, queue.IsExists(data1), true)
	queue.Denqueue()
	assert.EqualValues(t, queue.IsExists(data1), false)
}
