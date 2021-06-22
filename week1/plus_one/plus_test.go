package plusone

import (
	"testing"

	"gotest.tools/assert"
)

func TestPlusOne(t *testing.T) {
	assert.DeepEqual(t, PlusOne([]int{1, 2, 3}), []int{1, 2, 4})
	assert.DeepEqual(t, PlusOne([]int{4, 3, 2, 1}), []int{4, 3, 2, 2})
	assert.DeepEqual(t, PlusOne([]int{0}), []int{1})
	assert.DeepEqual(t, PlusOne([]int{9}), []int{1, 0})
	assert.DeepEqual(t, PlusOne([]int{9, 9, 9}), []int{1, 0, 0, 0})
	assert.DeepEqual(t, PlusOne([]int{7, 9, 9, 9}), []int{8, 0, 0, 0})
}
