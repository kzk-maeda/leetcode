package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	k := 0
	dummy := &ListNode{}
	tail := dummy

	for l1 != nil || l2 != nil || k > 0 {
		sum := k
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}
		if sum >= 10 {
			k = 1
			sum -= 10
		} else {
			k = 0
		}
		tail.Next = &ListNode{Val: sum}
		tail = tail.Next
	}

	return dummy.Next
}

func main() {

}
