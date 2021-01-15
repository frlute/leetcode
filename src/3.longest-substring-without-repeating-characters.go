/*
 * @lc app=leetcode id=3 lang=golang
 *
 * [3] Longest Substring Without Repeating Characters
 */

// @lc code=start
func lengthOfLongestSubstring(s string) int {
	windows := make(map[byte]int)
	left, right := 0, 0
	// 返回子字符串开始的索引
	res := 0

	for right < len(s) {
		str := s[right]
		right++
		// 进行窗口内数据的一系列更新
		windows[str]++
		// 判断左侧窗口是否要收缩
		for windows[str] > 1 {
			d := s[left]
			left++
			// 进行窗口内数据的一系列更新
			windows[d]--
		}
		// 在这里更新答案
		if res < (right - left) {
			res = right - left
		}
	}

	return res
}

// @lc code=end

