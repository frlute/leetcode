/*
 * @lc app=leetcode id=66 lang=golang
 *
 * [66] Plus One
 */

// @lc code=start
func plusOne(digits []int) []int {
	length := len(digits)
	for i:=length-1; i>=0; i-- {
		digits[i]++
		digits[i] = digits[i] % 10;
		if digits[i] != 0 {
			return digits
		}
	}
	
	return append([]int{1}, digits...)
}
// @lc code=end

