/*
 * @lc app=leetcode id=344 lang=golang
 *
 * [344] Reverse String
 */

// @lc code=start
func reverseString(s []byte)  {
	maxLen := len(s)
	for i:=0; i < maxLen/2; i++ {
		s[i], s[maxLen-i-1] = s[maxLen -i -1], s[i]
	}
}
// @lc code=end

