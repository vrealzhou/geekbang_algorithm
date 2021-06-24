package test

import (
	"strconv"
	"strings"
)

func ParseStringArray(v string) []string {
	stack := make([]bool, 0)
	result := make([]string, 0)
	buf := strings.Builder{}
	for _, c := range v {
		switch c {
		case '[':
			stack = append(stack, true)
			if len(stack) > 1 {
				buf.WriteRune(c)
			}
		case ']':
			if len(stack) > 1 {
				buf.WriteRune(c)
			} else {
				if buf.Len() > 0 {
					result = append(result, buf.String())
				}
			}
			stack = stack[:len(stack)-1]
		case ',':
			if len(stack) > 1 {
				buf.WriteRune(c)
			} else {
				if buf.Len() > 0 {
					result = append(result, buf.String())
				}
				buf = strings.Builder{}
			}
		default:
			buf.WriteRune(c)
		}
	}
	return result
}

func MsgAndArgs(i int, oprations, values []string) []interface{} {
	return []interface{}{"at %d, operation %s, value %s", i, oprations[i], values[i]}
}

func ExtractIntValue(item string, i int) int {
	v, err := strconv.Atoi(ParseStringArray(item)[i])
	if err != nil {
		panic(err)
	}
	return v
}
