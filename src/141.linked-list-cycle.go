/*
 * @lc app=leetcode id=141 lang=golang
 *
 * [141] Linked List Cycle
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}
	// 哈希法
	// cache := make(map[*ListNode]struct{})
	// for head.Next != nil {
	// 	if _, ok := cache[head]; ok {
	// 		return true
	// 	}
	// 	cache[head] = struct{}{}
	// 	head = head.Next
	// }
	// return false;

	// 快慢指针
	fast := head.Next
	for fast != nil && fast.Next != nil {
		head = head.Next
		fast = fast.Next.Next

		if head == fast {
			return true
		}
	}
	return false
}

// @lc code=end

