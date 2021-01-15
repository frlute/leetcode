/*
 * @lc app=leetcode id=438 lang=golang
 *
 * [438] Find All Anagrams in a String
 */

// @lc code=start
func findAnagrams(s string, p string) []int {
	needs := make(map[byte]int, 26)
	for i := range p {
		needs[p[i]]++
	}
	windows := make(map[byte]int)

	left, right := 0, 0
	res := []int{}
	match := 0
	for right < len(s) {
		str1 := s[right]
		right++

		if needs[str1] > 0 {
			windows[str1]++
			if windows[str1] == needs[str1] {
				match++
			}
		}

		// 收缩条件
		for right-left >= len(p) {
			// 在这里判断是否找到了合法的子串
			if match == len(needs) {
				res = append(res, left)
			}
			str2 := s[left]
			left++

			if needs[str2] > 0 {
				if windows[str2] == needs[str2] {
					match--
				}
				windows[str2]--
			}
		}
	}

	return res
}

// @lc code=end

