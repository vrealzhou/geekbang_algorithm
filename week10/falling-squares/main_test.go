package main

import (
	"testing"

	"gotest.tools/assert"
)

func TestRemoveDup(t *testing.T) {
	v := []int{1, 1, 2, 3, 3, 4}
	v = removeDup(v)
	t.Log(v)
}

func TestMain(t *testing.T) {
	var v []int
	// v = fallingSquares([][]int{{6, 1}, {9, 2}, {2, 4}})
	// assert.DeepEqual(t, v, []int{1, 2, 4})

	// v = fallingSquares([][]int{{1, 1}, {2, 2}})
	// assert.DeepEqual(t, v, []int{1, 2})

	// v = fallingSquares([][]int{{1, 5}, {2, 2}, {7, 5}})
	// assert.DeepEqual(t, v, []int{5, 7, 7})

	// v = fallingSquares([][]int{{1, 1}, {2, 1}})
	// assert.DeepEqual(t, v, []int{1, 1})

	v = fallingSquares([][]int{{9, 10}, {4, 1}, {2, 1}, {7, 4}, {6, 10}})
	assert.DeepEqual(t, v, []int{10, 10, 10, 14, 24})
}
