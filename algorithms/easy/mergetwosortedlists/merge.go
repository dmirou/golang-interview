package mergetwosortedlists

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil && list2 == nil {
		return nil
	}

	if list1 == nil {
		return list2
	}

	if list2 == nil {
		return list1
	}

	var h *ListNode

	if list1.Val <= list2.Val {
		h = list1
		list1 = list1.Next
	} else {
		h = list2
		list2 = list2.Next
	}

	cur := h

	for {
		if list1 == nil {
			cur.Next = list2
			break
		}
		if list2 == nil {
			cur.Next = list1
			break
		}
		if list1.Val <= list2.Val {
			cur.Next = list1
			cur = list1
			list1 = list1.Next
			continue
		}
		cur.Next = list2
		cur = list2
		list2 = list2.Next
	}

	return h
}
