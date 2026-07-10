/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reverseList(head *ListNode) *ListNode {
	
	if head ==nil{
		return head
	}
	
	var curr *ListNode

	for head.Next != nil{
		temp := head.Next
		head.Next = curr
		curr = head
		head = temp
	}

	head.Next = curr

	return head
    
}
