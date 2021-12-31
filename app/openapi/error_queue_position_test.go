package openapi_test

import (
	"testing"

	openapi "github.com/hazuki3417/testing-file-generator/openapi"
	"github.com/stretchr/testify/assert"
)

func Test_ErrorPositionQueue_Enqueue(t *testing.T) {
	queue := openapi.NewErrorPositionQueue(5)

	assert.EqualValues(t, queue.IsEmpty(), true)
	assert.EqualValues(t, queue.IsNotEmpty(), false)
	assert.EqualValues(t, queue.HasEnqueue(), true)
	assert.EqualValues(t, queue.HasDenqueue(), false)
	assert.EqualValues(t, queue.Size(), 0)

	err1 := openapi.ErrorPosition{Key: "name", Reason: "not empty"}
	queue.Enqueue(err1)

	assert.EqualValues(t, queue.IsEmpty(), false)
	assert.EqualValues(t, queue.IsNotEmpty(), true)
	assert.EqualValues(t, queue.HasEnqueue(), false)
	assert.EqualValues(t, queue.HasDenqueue(), true)
	assert.EqualValues(t, queue.Size(), 1)

	err2 := openapi.ErrorPosition{Key: "age", Reason: "not empty"}
	queue.Enqueue(err2)

	assert.EqualValues(t, queue.IsEmpty(), false)
	assert.EqualValues(t, queue.IsNotEmpty(), true)
	assert.EqualValues(t, queue.HasEnqueue(), false)
	assert.EqualValues(t, queue.HasDenqueue(), true)
	assert.EqualValues(t, queue.Size(), 2)

	assert.EqualValues(t, err1, queue.Denqueue())
	assert.EqualValues(t, queue.IsEmpty(), false)
	assert.EqualValues(t, queue.IsNotEmpty(), true)
	assert.EqualValues(t, queue.HasEnqueue(), false)
	assert.EqualValues(t, queue.HasDenqueue(), true)
	assert.EqualValues(t, queue.Size(), 1)

	assert.EqualValues(t, err2, queue.Denqueue())
	assert.EqualValues(t, queue.IsEmpty(), true)
	assert.EqualValues(t, queue.IsNotEmpty(), false)
	assert.EqualValues(t, queue.HasEnqueue(), true)
	assert.EqualValues(t, queue.HasDenqueue(), false)
	assert.EqualValues(t, queue.Size(), 0)

	data1 := openapi.ErrorPosition{Key: "key1", Reason: "reason1"}
	queue.Enqueue(data1)
	data2 := openapi.ErrorPosition{Key: "key2", Reason: "reason2"}
	queue.Enqueue(data2)
	data3 := openapi.ErrorPosition{Key: "key3", Reason: "reason3"}
	queue.Enqueue(data3)

	assert.EqualValues(t, queue.Size(), 3)
	assert.EqualValues(t, queue.All(), []openapi.ErrorPosition{data1, data2, data3})
	assert.EqualValues(t, queue.Size(), 0)

	queue.Enqueue(openapi.ErrorPosition{Key: "key0", Reason: "reason0"})
	queue.Denqueue()

}
