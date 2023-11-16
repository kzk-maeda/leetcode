package main

import (
	"fmt"
	"sort"
)

func removeElement(nums []int, val int) int {
	k := 0
	for _, v := range nums {
		if v != val {
			nums[k] = v
			k++
		}
	}
	fmt.Println(nums)
	return k
}

func removeElementEfficiency(nums []int, val int) int {
	result := len(nums)

	for i, num := range nums {
		if num == val {
			nums[i] = 101
			result--
		}
	}
	fmt.Println(nums)
	sort.Ints(nums)
	fmt.Println(nums)
	return result
}

func main() {
	nums := []int{3, 2, 2, 3}
	val := 3
	k := removeElementEfficiency(nums, val)
	println(k)
}
