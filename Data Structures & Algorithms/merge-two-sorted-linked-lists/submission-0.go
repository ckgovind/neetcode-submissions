
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {

	if list1 == nil{
		return list2
	}

	if list2 == nil{
		return list1
	}


	if list1.Val < list2.Val {
		next := list1.Next
		list1.Next = mergeTwoLists(next, list2)
		return list1
	} else {
		next := list2.Next
		list2.Next = mergeTwoLists(list1, next)
		return list2
	}

}
