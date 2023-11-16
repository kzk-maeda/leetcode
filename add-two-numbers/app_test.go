package main

import "testing"

type TestVals struct {
	l1       *ListNode
	l2       *ListNode
	expected *ListNode
}

func createList(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}

	head := &ListNode{Val: nums[0]}
	current := head
	for _, num := range nums[1:] {
		current.Next = &ListNode{Val: num}
		current = current.Next
	}
	return head
}

func Test(t *testing.T) {
	tests := []TestVals{
		{
			l1:       createList([]int{2, 4, 3}),
			l2:       createList([]int{5, 6, 4}),
			expected: createList([]int{7, 0, 8}),
		},
		{
			l1:       createList([]int{9, 9, 9, 9, 9, 9, 9}),
			l2:       createList([]int{9, 9, 9, 9}),
			expected: createList([]int{8, 9, 9, 9, 0, 0, 0, 1}),
		},
	}

	for _, test := range tests {
		actual := addTwoNumbers(test.l1, test.l2) // AddTwoNumbers は解決すべき関数
		for a, e := actual, test.expected; a != nil || e != nil; a, e = a.Next, e.Next {
			if a == nil || e == nil || a.Val != e.Val {
				t.Errorf("addTwoNumbers(%v, %v) = %v, want %v", test.l1, test.l2, actual, test.expected)
				break
			}
		}
	}
}
