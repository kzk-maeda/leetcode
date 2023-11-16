package main

import "fmt"

func romanToInt(s string) int {
	romanValues := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	total := 0
	n := len(s)
	for i := 0; i < n; i++ {
		value := romanValues[s[i]]
		if i+1 < n && value < romanValues[s[i+1]] {
			total -= value
		} else {
			total += value
		}
	}
	return total
}

func main() {
	result := romanToInt("MCMXCIV")
	fmt.Println(result)
}
