/*
 * @lc app=leetcode id=19 lang=golang
 *
 * [19] Remove Nth Node From End of List
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	// result := &ListNode{}
	// result.Next = head

	// var pre *ListNode
	// cur := result
	// i := 1
	// for head != nil {
	// 	if i >= n {
	// 		pre = cur
	// 		cur = cur.Next
	// 	}
	// 	head = head.Next
	// 	i++
	// }

	// pre.Next = pre.Next.Next
	// return result.Next

	// 快慢指针
	fast, slow := head, head
	// 快指针先前进 n 步
	for ; n > 0; n-- {
		fast = fast.Next
	}

	// 如果此时快指针走到头了，
	// 说明倒数第 n 个节点就是第一个节点
	if fast == nil {
		return head.Next
	}

	// 让慢指针和快指针同步向前
	for fast != nil && fast.Next != nil {
		fast = fast.Next
		slow = slow.Next
	}

	// slow.Next 就是倒数第 n 个节点，删除它
	slow.Next = slow.Next.Next

	return head
}

// @lc code=end

