/*
 * @lc app=leetcode id=142 lang=golang
 *
 * [142] Linked List Cycle II
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func detectCycle(head *ListNode) *ListNode {
	// if head == nil || head.Next == nil {
	// 	return nil
	// }
	fast, slow := head, head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow {
			break
		}
	}
	// 遇到空指针说明无环
	if fast == nil || fast.Next == nil {
		return nil
	}

	for slow = head; slow != fast; fast, slow = fast.Next, slow.Next {
	}

	return slow
}

// @lc code=end

