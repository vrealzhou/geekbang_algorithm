package deque

import (
	"fmt"
	"testing"

	"gotest.tools/assert"
)

func TestDeque(t *testing.T) {
	q := NewArrayDeque(3)
	assert.Equal(t, q.IsEmpty(), true)
	assert.Equal(t, q.IsFull(), false)
	assert.Equal(t, q.DeleteLast(), false)
	assert.Equal(t, q.DeleteFront(), false)

	assert.Equal(t, q.InsertLast(1), true)
	fmt.Println(q.String())
	assert.Equal(t, q.GetFront(), 1)
	assert.Equal(t, q.GetRear(), 1)
	assert.Equal(t, q.IsEmpty(), false)
	assert.Equal(t, q.IsFull(), false)

	assert.Equal(t, q.InsertLast(2), true)
	fmt.Println(q.String())
	assert.Equal(t, q.GetFront(), 1)
	assert.Equal(t, q.GetRear(), 2)
	assert.Equal(t, q.IsEmpty(), false)
	assert.Equal(t, q.IsFull(), false)

	assert.Equal(t, q.InsertFront(3), true)
	assert.Equal(t, q.GetFront(), 3)
	assert.Equal(t, q.GetRear(), 2)
	assert.Equal(t, q.IsEmpty(), false)
	assert.Equal(t, q.IsFull(), true)

	assert.Equal(t, q.InsertFront(4), false)
	assert.Equal(t, q.IsEmpty(), false)
	assert.Equal(t, q.IsFull(), true)
	assert.Equal(t, q.GetFront(), 3)
	assert.Equal(t, q.GetRear(), 2)
	assert.Equal(t, q.IsFull(), true)

	assert.Equal(t, q.DeleteLast(), true)
	assert.Equal(t, q.GetFront(), 3)
	assert.Equal(t, q.GetRear(), 1)
	assert.Equal(t, q.IsEmpty(), false)
	assert.Equal(t, q.IsFull(), false)

	assert.Equal(t, q.InsertFront(4), true)
	assert.Equal(t, q.GetFront(), 4)
	assert.Equal(t, q.GetRear(), 1)
	assert.Equal(t, q.IsEmpty(), false)
	assert.Equal(t, q.IsFull(), true)

	assert.Equal(t, q.DeleteFront(), true)
	assert.Equal(t, q.GetFront(), 3)
	assert.Equal(t, q.GetRear(), 1)
	assert.Equal(t, q.IsEmpty(), false)
	assert.Equal(t, q.IsFull(), false)

	assert.Equal(t, q.DeleteFront(), true)
	assert.Equal(t, q.GetFront(), 1)
	assert.Equal(t, q.GetRear(), 1)
	assert.Equal(t, q.IsEmpty(), false)
	assert.Equal(t, q.IsFull(), false)

	assert.Equal(t, q.DeleteLast(), true)
	assert.Equal(t, q.GetFront(), -1)
	assert.Equal(t, q.GetRear(), -1)
	assert.Equal(t, q.IsEmpty(), true)
	assert.Equal(t, q.IsFull(), false)
}

func TestDeque2(t *testing.T) {
	q := NewArrayDeque(3)
	assert.Equal(t, q.InsertFront(8), true)
	assert.Equal(t, q.InsertLast(8), true)
	assert.Equal(t, q.InsertLast(2), true)
	assert.Equal(t, q.GetFront(), 8)
	assert.Equal(t, q.DeleteLast(), true)
	assert.Equal(t, q.GetRear(), 8)
	assert.Equal(t, q.InsertFront(9), true)
	assert.Equal(t, q.DeleteFront(), true)
	assert.Equal(t, q.GetRear(), 8)
	assert.Equal(t, q.InsertLast(2), true)
	fmt.Println(q.String())
	assert.Equal(t, q.IsFull(), true)
}
