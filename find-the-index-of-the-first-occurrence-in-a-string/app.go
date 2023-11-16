package main

import "fmt"

func strStr(haystack string, needle string) int {
	for i := 0; i <= len(haystack)-len(needle); i++ {
		if haystack[i:len(needle)+i] == needle {
			return i
		}
	}
	return -1
}
func main() {
	// haystack := "sadbutsad"
	// needle := "sad"
	// k := strStr(haystack, needle)
	// fmt.Println(k)

	// fmt.Println(strStr("mississippi", "issipi"))
	fmt.Println(strStr("abc", "c"))
}
