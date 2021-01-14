/*
 * @lc app=leetcode id=349 lang=golang
 *
 * [349] Intersection of Two Arrays
 */

// @lc code=start
func intersection(nums1 []int, nums2 []int) []int {
	if len(nums1) == 0 || len(nums2) == 0 {
		return []int{}
	}

	m1 := make(map[int]bool, len(nums1)) 
	for _, value := range nums1 {
		m1[value] = true
	}

	res := make([]int, 0)
	for _, value := range nums2 {
		if m1[value] == true {
			res = append(res, value)
			m1[value] = false
		}
	}
	return res
}
// @lc code=end

