/*
 * @lc app=leetcode id=14 lang=golang
 *
 * [14] Longest Common Prefix
 */
// @lc code=start
func longestCommonPrefix1(strs []string) string {
	if len(strs) < 1 {
		return ""
	}

	prefix := strs[0]
	for _, str := range strs {
		for strings.Index(str, prefix) != 0 {
			if prefix == "" {
				return ""
			}
			prefix = prefix[:len(prefix) - 1]
		}
	}

	return prefix
}
// @lc code=end

