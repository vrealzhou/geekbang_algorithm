package test

import (
	"testing"

	"gotest.tools/assert"
)

func TestParseArray(t *testing.T) {
	assert.DeepEqual(t, ParseStringArray(`["MyCircularDeque","insertFront","getRear","deleteLast"]`), []string{`"MyCircularDeque"`, `"insertFront"`, `"getRear"`, `"deleteLast"`})
	assert.DeepEqual(t, ParseStringArray(`[[77],[89],[],[],[],[19]]`), []string{"[77]", "[89]", "[]", "[]", "[]", "[19]"})
	assert.DeepEqual(t, ParseStringArray(`[]`), []string{})
}
