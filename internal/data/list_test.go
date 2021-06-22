package data

import (
	"testing"

	"gotest.tools/assert"
	is "gotest.tools/assert/cmp"
)

func TestNewIntList(t *testing.T) {
	t.Run("empty", func(t *testing.T) {
		n := IntList()
		assert.Assert(t, is.Nil(n))
	})
	t.Run("1,2,3", func(t *testing.T) {
		n := IntList(1, 2, 3)
		assert.Equal(t, n.Val, 1)
		n = n.Next
		assert.Equal(t, n.Val, 2)
		n = n.Next
		assert.Equal(t, n.Val, 3)
		n = n.Next
		assert.Assert(t, is.Nil(n))
	})
}

func TestNewIntListToSlice(t *testing.T) {
	t.Run("nil", func(t *testing.T) {
		var n *IntListNode
		assert.DeepEqual(t, n.Slice(), []int{})
	})
	t.Run("empty", func(t *testing.T) {
		n := &IntListNode{}
		assert.DeepEqual(t, n.Slice(), []int{0})
	})
	t.Run("1,2,3", func(t *testing.T) {
		n := IntList(1, 2, 3)
		assert.DeepEqual(t, n.Slice(), []int{1, 2, 3})
	})
}
