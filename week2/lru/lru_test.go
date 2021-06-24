package lru

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/vrealzhou/geekbang_algorithm/internal/test"
	"gotest.tools/assert"
)

func TestLRU(t *testing.T) {
	testSequence(t,
		`["LRUCache","put","put","get","put","get","put","get","get","get"]`,
		`[[2],[1,1],[2,2],[1],[3,3],[2],[4,4],[1],[3],[4]]`,
		`[null,null,null,1,null,-1,null,-1,3,4]`,
	)
	testSequence(t,
		`["LRUCache","put","put","put","put","get","get"]`,
		`[[2],[2,1],[1,1],[2,3],[4,1],[1],[2]]`,
		`[null,null,null,null,null,-1,3]`,
	)
}

func testSequence(t *testing.T, oprationStr, valueStr, expectStr string) {
	oprations := test.ParseStringArray(oprationStr)
	values := test.ParseStringArray(valueStr)
	expects := test.ParseStringArray(expectStr)
	t.Run(fmt.Sprintf("case %s %s", oprationStr, valueStr), func(t *testing.T) {
		var d *LRUCache
		for i, opration := range oprations {
			switch opration {
			case `"LRUCache"`:
				tmp := Constructor(test.ExtractIntValue(values[i], 0))
				d = &tmp
			case `"put"`:
				d.Put(test.ExtractIntValue(values[i], 0), test.ExtractIntValue(values[i], 1))
			case `"get"`:
				result := d.Get(test.ExtractIntValue(values[i], 0))
				assert.Equal(t, strconv.Itoa(result), expects[i], test.MsgAndArgs(i, oprations, values)...)
			}
		}
	})
}
