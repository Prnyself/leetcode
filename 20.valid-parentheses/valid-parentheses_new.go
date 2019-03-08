package main

import (
	"fmt"
	"strconv"
)

type charStringStack struct {
	str []int32
}

func (stack *charStringStack) pop() int32{
	n := len(stack.str)
	if n == 0 {
		return ''
	}
	poped := stack.str[n-1]
	stack.str = stack.str[0:n-1]
	return poped
}

func (stack *charStringStack) push(val int32) {
	stack.str = append(stack.str, val)
}

func main() {
	parenthes := map[int32]int32{
		'(': ')',
		'[': ']',
		'{': '}',
	}
	fmt.Println(parenthes)
	strconv.Atoi()
}

func isValidNew(s string) bool {
	parenthes := map[int32]int32{
		'(': ')',
		'[': ']',
		'{': '}',
	}

	str := &charStringStack{
		str: []int32{},
	}

	for _, v := range s {
		if _, ok := parenthes[v]; ok {
			str.push(v)
		} else {
			if tmp, ok := parenthes[str.pop()]; !ok || tmp != v {
				return false
			}
		}
	}
	return len(str.str) == 0
}