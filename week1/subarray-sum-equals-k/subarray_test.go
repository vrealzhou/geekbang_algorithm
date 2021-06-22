package subarray

import (
	"testing"

	"gotest.tools/assert"
)

func TestSubarraySum(t *testing.T) {
	assert.Equal(t, subarraySum([]int{1, 1, 1}, 2), 2)
	assert.Equal(t, subarraySum([]int{1, 1, -1, 1, 5, -3, 0, 2, 1}, 3), 6)
	assert.Equal(t, subarraySum([]int{1, 1, 1, 1, 5, 3, 0, 2, 1}, 3), 6)
	assert.Equal(t, subarraySum([]int{1, 1, 1, 1, -1, 1, -1, 1, -1}, 3), 7)
	assert.Equal(t, subarraySum([]int{1}, 1), 1)
	assert.Equal(t, subarraySum([]int{1, 1}, 1), 2)
	assert.Equal(t, subarraySum([]int{0, 1}, 1), 2)
	assert.Equal(t, subarraySum([]int{0, 1}, 2), 0)
}
