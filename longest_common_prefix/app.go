package main

import (
	"fmt"
	"strings"
)

func longestCommonPrefix(strs []string) string {
	var pref string
	if len(strs) == 0 {
		return pref
	}
	if len(strs) == 1 {
		return strs[0]
	}
	for i := 0; i < len(strs[0]); i++ {
		for j := range strs {
			if i >= len(strs[j]) || strs[0][i] != strs[j][i] {
				return pref
			}
		}
		pref += string(strs[0][i])
	}
	return pref
}

func longestCommonPrefix2(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	prefix := strs[0]
	for i := 1; i < len(strs); i++ {
		for !strings.HasPrefix(strs[i], prefix) {
			prefix = prefix[:len(prefix)-1]
			if len(prefix) == 0 {
				return ""
			}
		}
	}
	return prefix
}

func main() {
	// strs := []string{"flower", "flow", "flight"}
	// pref := longestCommonPrefix(strs)
	// fmt.Println(pref)

	// strs = []string{"dog", "racecar", "car"}
	// pref = longestCommonPrefix(strs)
	// fmt.Println(pref)

	// strs = []string{"a"}
	// pref = longestCommonPrefix(strs)
	// fmt.Println(pref)

	strs := []string{"flower", "flower", "flower", "flower"}
	fmt.Println(longestCommonPrefix2(strs))
}
