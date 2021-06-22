package merge

import (
	"testing"

	"gotest.tools/assert"

	"github.com/vrealzhou/geekbang_algorithm/internal/data"
)

func TestMergeTwoLists(t *testing.T) {
	assert.DeepEqual(t, mergeTwoLists(data.IntList(1, 2, 4), data.IntList(1, 3, 4)).Slice(), []int{1, 1, 2, 3, 4, 4})
	assert.DeepEqual(t, mergeTwoLists(data.IntList(), data.IntList()).Slice(), []int{})
	assert.DeepEqual(t, mergeTwoLists(data.IntList(), data.IntList(0)).Slice(), []int{0})
}
