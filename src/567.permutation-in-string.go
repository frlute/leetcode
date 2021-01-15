/*
 * @lc app=leetcode id=567 lang=golang
 *
 * [567] Permutation in String
 */

// @lc code=start
func checkInclusion(s1 string, s2 string) bool {
	if len(s2) > len(s1) {
		return false
	}
	needs := make(map[byte]int, len(s2))
	for _, value := range []byte(s2) {
		needs[value]++
	}
	windows := make(map[byte]int, len(s1))

	left, right := 0, 0
	match := 0
	for right < len(s1) {
		addStr := s1[right]
		right++

		if needs[addStr] > 0 {
			windows[addStr]++
			if windows[addStr] == needs[addStr] {
				match++
			}
		}

		// 收缩条件
		for right-left >= len(s2) {
			// 在这里判断是否找到了合法的子串
			if match == len(needs) {
				return true
			}
			delStr := s1[left]
			left--

			if needs[delStr] > 0 {
				windows[delStr]--
				if windows[delStr] < needs[delStr] {
					match--
				}
			}
		}
	}

	return false
}

// @lc code=end

