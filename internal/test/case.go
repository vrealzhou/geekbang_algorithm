package test

import "strings"

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
