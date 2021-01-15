/*
 * @lc app=leetcode id=76 lang=golang
 *
 * [76] Minimum Window Substring
 */

// @lc code=start
func minWindow(s string, t string) string {
	if len(s) < len(t) {
		return ""
	}
	needs := make(map[byte]int, len(t))
	windows := make(map[byte]int, len(s))

	for _, c := range []byte(t) {
		needs[c]++
	}

	left, right := 0, 0
	// 记录 window 中已经有多少字符符合要求了
	match := 0
	// 记录最小覆盖子串的起始索引及长度
	start, minLen := 0, math.MaxInt32
	for right < len(s) {
		// c 是将移入窗口的字符
		c := s[right]
		// 右移窗口
		right++

		// 进行窗口内数据的一系列更新
		if _, ok := needs[c]; ok {
			windows[c]++
			if windows[c] == needs[c] {
				// 字符 c 的出现次数符合要求了
				match++
			}
		}

		// 判断左侧窗口是否要收缩
		for match == len(needs) {
			// 在这里更新最小覆盖子串
			if right-left < minLen {
				// 更新最小子串的位置和长度
				start = left
				minLen = right - left
			}
			// d 是将移出窗口的字符
			d := s[left]
			// 左移窗口
			left++

			// 进行窗口内数据的一系列更新
			if _, ok := needs[d]; ok {
				// 移出 windows
				windows[d]--
				if windows[d] < needs[d] {
					// 字符 d 出现次数不再符合要求了
					match--
				}
			}
		}

	}
	if minLen == math.MaxInt32 {
		return ""
	}
	return s[start : start+minLen]
}

// @lc code=end

