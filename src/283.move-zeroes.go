/*
 * @lc app=leetcode id=283 lang=golang
 *
 * [283] Move Zeroes
 */

// @lc code=start
func moveZeroes(nums []int)  {
	// 双指针法
    if len(nums) <= 1 {
		return
	}

	moveEle := 0
	slow := 0;
	for i := 0; i < len(nums); i++ {
		if nums[i] != moveEle {
			nums[i], nums[slow] = nums[slow], nums[i]
			slow++
		}
	}
}
// @lc code=end

