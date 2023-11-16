package main

import (
	"fmt"
)

func main() {
	x := 12121
	var result bool = false
	if x < 0 || (x != 0 && x%10 == 0) { // 負の数、0、10の倍数は回文になり得ない
		result = false
	}

	var reversed int
	for x > reversed {
		reversed = reversed*10 + x%10
		x /= 10
	}

	result = x == reversed || x == reversed/10
	fmt.Println(result)
}
