/*
 * @lc app=leetcode id=567 lang=golang
 *
 * [567] Permutation in String
 */

// @lc code=start
func checkInclusion(s1 string, s2 string) bool {
	if len(s1) > len(s2) {
		return false
	}
	needs := make(map[byte]int, 26)
	for _, value := range []byte(s1) {
		needs[value]++
	}
	windows := make(map[byte]int, 26)

	left, right := 0, 0
	match := 0
	for right < len(s2) {
		addStr := s2[right]
		right++

		if needs[addStr] > 0 {
			windows[addStr]++
			if windows[addStr] == needs[addStr] {
				match++
			}
		}

		// 收缩条件
		for right-left >= len(s1) {
			// 在这里判断是否找到了合法的子串
			if match == len(needs) {
				return true
			}
			delStr := s2[left]
			left++

			if needs[delStr] > 0 {
				// windows[delStr]--
				// if windows[delStr] < needs[delStr] {
				// 	match--
				// }
				if windows[delStr] == needs[delStr] {
					match--
				}
				windows[delStr]--
			}
		}
	}

	return false
}

// @lc code=end

