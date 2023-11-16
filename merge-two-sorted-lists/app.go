package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	// ダミーヘッドを作成します。
	dummy := &ListNode{}
	tail := dummy

	// 両方のリストにノードがある間、小さい方のノードを新しいリストに追加します。
	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			tail.Next = list1
			list1 = list1.Next
		} else {
			tail.Next = list2
			list2 = list2.Next
		}
		tail = tail.Next
	}

	// いずれかのリストが終わった場合、残りのリストを追加します。
	if list1 != nil {
		tail.Next = list1
	} else {
		tail.Next = list2
	}

	// ダミーヘッドの次のノードを返します。
	return dummy.Next
}

func main() {
	print(mergeTwoLists(&ListNode{1, &ListNode{2, &ListNode{4, nil}}}, &ListNode{1, &ListNode{3, &ListNode{4, nil}}}))
}
