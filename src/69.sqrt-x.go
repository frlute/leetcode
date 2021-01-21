/*
 * @lc app=leetcode id=69 lang=golang
 *
 * [69] Sqrt(x)
 */

// @lc code=start
func mySqrt(x int) int {
	if x < 2 {
		return x
	}

	left, right := 1, x/2
	for left <= right {
		mid := left + ((right - left) >> 1)
		if mid == x/mid {
			return mid
		} else if mid < x/mid {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return right
}

// @lc code=end

