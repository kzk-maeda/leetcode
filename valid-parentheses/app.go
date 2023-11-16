package main

import (
	"fmt"
)

func isValid(s string) bool {
	if len(s)%2 != 0 {
		return false
	}
	if s[0] == ')' || s[0] == ']' || s[0] == '}' {
		return false
	}

	parents := map[rune]rune{
		'(': ')',
		'[': ']',
		'{': '}',
	}
	var stack []rune

	for _, c := range s {
		if closing, exists := parents[c]; exists {
			stack = append(stack, closing)
		} else {
			if len(stack) == 0 {
				return false
			}
			if stack[len(stack)-1] != c { // if closing bracket doesn't match last opening bracket
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}
	return len(stack) == 0
}

func main() {
	fmt.Println(isValid("(("))
	fmt.Println(isValid("()[]{}")) // true
	fmt.Println(isValid("(]"))     // false
	fmt.Println(isValid("(){}}{")) // false
	fmt.Println(isValid("{[]}"))   // true
}
