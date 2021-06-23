package deque

import (
	"testing"
)

func TestArrayDeque(t *testing.T) {
	testCases(t, func(k int) Deque {
		return NewArrayDeque(k)
	})
}
