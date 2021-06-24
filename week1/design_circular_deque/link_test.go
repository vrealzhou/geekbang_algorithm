package deque

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/vrealzhou/geekbang_algorithm/internal/test"
	"gotest.tools/assert"
)

type Constructor func(k int) Deque

type Deque interface {
	InsertFront(value int) bool
	InsertLast(value int) bool
	DeleteFront() bool
	DeleteLast() bool
	GetFront() int
	GetRear() int
	IsEmpty() bool
	IsFull() bool
}

func testCases(t *testing.T, constructor func(k int) Deque) {
	testSequence(t, constructor,
		`["MyCircularDeque","isEmpty","isFull","deleteLast","deleteFront","insertLast","getFront","getRear","isEmpty","isFull","insertLast","getFront","getRear","isEmpty","isFull","insertFront","getFront","getRear","isEmpty","isFull","insertFront","isEmpty","isFull","getFront","getRear","isFull","deleteLast","getFront","getRear","isEmpty","isFull","insertFront","getFront","getRear","isEmpty","isFull","deleteFront","getFront","getRear","isEmpty","isFull","deleteFront","getFront","getRear","isEmpty","isFull","deleteLast","getFront","getRear","isEmpty","isFull"]`,
		`[[3],[],[],[],[],[1],[],[],[],[],[2],[],[],[],[],[3],[],[],[],[],[4],[],[],[],[],[],[],[],[],[],[],[4],[],[],[],[],[],[],[],[],[],[],[],[],[],[],[],[],[],[],[]]`,
		`[null,true,false,false,false,true,1,1,false,false,true,1,2,false,false,true,3,2,false,true,false,false,true,3,2,true,true,3,1,false,false,true,4,1,false,true,true,3,1,false,false,true,1,1,false,false,true,-1,-1,true,false]`,
	)
	testSequence(t, constructor,
		`["MyCircularDeque","insertFront","getRear","deleteLast","getRear","insertFront","insertFront","insertFront","insertFront","isFull","insertFront","isFull","getRear","deleteLast","getFront","getFront","insertLast","deleteFront","getFront","insertLast","getRear","insertLast","getRear","getFront","getFront","getFront","getRear","getRear","insertFront","getFront","getFront","getFront","getFront","deleteFront","insertFront","getFront","deleteLast","insertLast","insertLast","getRear","getRear","getRear","isEmpty","insertFront","deleteLast","getFront","deleteLast","getRear","getFront","isFull","isFull","deleteFront","getFront","deleteLast","getRear","insertFront","getFront","insertFront","insertFront","getRear","isFull","getFront","getFront","insertFront","insertLast","getRear","getRear","deleteLast","insertFront","getRear","insertLast","getFront","getFront","getFront","getRear","insertFront","isEmpty","getFront","getFront","insertFront","deleteFront","insertFront","deleteLast","getFront","getRear","getFront","insertFront","getFront","deleteFront","insertFront","isEmpty","getRear","getRear","getRear","getRear","deleteFront","getRear","isEmpty","deleteFront","insertFront","insertLast","deleteLast"]`,
		`[[77],[89],[],[],[],[19],[23],[23],[82],[],[45],[],[],[],[],[],[74],[],[],[98],[],[99],[],[],[],[],[],[],[8],[],[],[],[],[],[75],[],[],[35],[59],[],[],[],[],[22],[],[],[],[],[],[],[],[],[],[],[],[21],[],[26],[63],[],[],[],[],[87],[76],[],[],[],[26],[],[67],[],[],[],[],[36],[],[],[],[72],[],[87],[],[],[],[],[85],[],[],[91],[],[],[],[],[],[],[],[],[],[34],[44],[]]`,
		`[null,true,89,true,-1,true,true,true,true,false,true,false,19,true,45,45,true,true,82,true,98,true,99,82,82,82,99,99,true,8,8,8,8,true,true,75,true,true,true,59,59,59,false,true,true,22,true,98,22,false,false,true,75,true,74,true,21,true,true,74,false,63,63,true,true,76,76,true,true,74,true,26,26,26,67,true,false,36,36,true,true,true,true,87,74,87,true,85,true,true,false,74,74,74,74,true,74,false,true,true,true,true]`,
	)
	testSequence(t, constructor,
		`["MyCircularDeque","insertFront","getFront","insertFront","getFront","deleteLast","insertFront","insertFront","getRear","getFront","getRear","getRear","insertLast","deleteFront","getFront","insertLast","getRear","insertLast","deleteFront","insertFront","isFull","getRear","deleteLast","insertLast","getRear","getFront","getFront","insertLast","insertFront","deleteFront","getRear","insertLast","deleteFront","insertFront","insertFront","getRear","getFront","insertFront","insertLast","getRear","getFront","insertFront","insertFront","insertLast","insertLast","getRear","isEmpty","deleteFront","getRear","getRear","getRear","insertLast","getFront","getFront","deleteLast","deleteLast","insertLast","getRear","getRear","insertLast","insertLast","insertFront","getFront","getRear","getFront","insertFront","insertFront","deleteFront","isEmpty","getFront","deleteFront","isFull","getFront","getRear","insertLast","getFront","insertLast","getRear","insertLast","insertFront","getRear","getFront","getFront","deleteLast","deleteLast","insertLast","getRear","getRear","getFront","deleteLast","isFull","insertLast","insertLast","insertFront","getFront","insertFront","isFull","getRear","insertFront","deleteLast","insertLast","insertLast"]`,
		`[[52],[80],[],[27],[],[],[60],[81],[],[],[],[],[46],[],[],[98],[],[11],[],[51],[],[],[],[28],[],[],[],[28],[69],[],[],[11],[],[25],[74],[],[],[48],[7],[],[],[65],[59],[23],[32],[],[],[],[],[],[],[84],[],[],[],[],[64],[],[],[17],[34],[46],[],[],[],[6],[20],[],[],[],[],[],[],[],[34],[],[66],[],[54],[34],[],[],[],[],[],[43],[],[],[],[],[],[21],[93],[79],[],[8],[],[],[78],[],[7],[67]]`,
		`[null,true,80,true,27,true,true,true,27,81,27,27,true,true,60,true,98,true,true,true,false,11,true,true,28,51,51,true,true,true,28,true,true,true,true,11,74,true,true,7,48,true,true,true,true,32,false,true,32,32,32,true,65,65,true,true,true,64,64,true,true,true,46,34,46,true,true,true,false,6,true,false,46,34,true,46,true,66,true,true,54,34,34,true,true,true,43,43,34,true,false,true,true,true,79,true,false,93,true,true,true,true]`,
	)
	testSequence(t, constructor,
		`["MyCircularDeque","insertFront","insertLast","insertLast",getFront","deleteLast","getRear","insertFront","deleteFront","getRear","insertLast","isFull"]`,
		`[[3],[8],[8],[2],[],[],[],[9],[],[],[2],[]]`,
		`[null,true,true,true,8,true,8,true,true,8,true,true]`,
	)
}

func TestListDeque(t *testing.T) {
	testCases(t, func(k int) Deque {
		return NewListDeque(k)
	})
}

func testSequence(t *testing.T, constructor func(k int) Deque, oprationStr, valueStr, expectStr string) {
	oprations := test.ParseStringArray(oprationStr)
	values := test.ParseStringArray(valueStr)
	expects := test.ParseStringArray(expectStr)
	t.Run(fmt.Sprintf("case %s %s", oprationStr, valueStr), func(t *testing.T) {
		var d Deque
		for i, opration := range oprations {
			switch opration {
			case `"MyCircularDeque"`:
				d = constructor(test.ExtractIntValue(values[i], 0))
			case `"insertFront"`:
				result := d.InsertFront(test.ExtractIntValue(values[i], 0))
				assert.Equal(t, strconv.FormatBool(result), expects[i], test.MsgAndArgs(i, oprations, values)...)
			case `"insertLast"`:
				result := d.InsertLast(test.ExtractIntValue(values[i], 0))
				assert.Equal(t, strconv.FormatBool(result), expects[i], test.MsgAndArgs(i, oprations, values)...)
			case `"deleteFront"`:
				result := d.DeleteFront()
				assert.Equal(t, strconv.FormatBool(result), expects[i], test.MsgAndArgs(i, oprations, values)...)
			case `"deleteLast"`:
				result := d.DeleteLast()
				assert.Equal(t, strconv.FormatBool(result), expects[i], test.MsgAndArgs(i, oprations, values)...)
			case `"getFront"`:
				result := d.GetFront()
				assert.Equal(t, strconv.Itoa(result), expects[i], test.MsgAndArgs(i, oprations, values)...)
			case `"getRear"`:
				result := d.GetRear()
				assert.Equal(t, strconv.Itoa(result), expects[i], test.MsgAndArgs(i, oprations, values)...)
			case `"isEmpty"`:
				result := d.IsEmpty()
				assert.Equal(t, strconv.FormatBool(result), expects[i], test.MsgAndArgs(i, oprations, values)...)
			case `"isFull"`:
				result := d.IsFull()
				assert.Equal(t, strconv.FormatBool(result), expects[i], test.MsgAndArgs(i, oprations, values)...)
			}
		}
	})
}
