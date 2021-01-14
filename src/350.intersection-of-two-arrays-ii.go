/*
 * @lc app=leetcode id=350 lang=golang
 *
 * [350] Intersection of Two Arrays II
 */

 func intersect1(nums1 []int, nums2 []int) []int {
	// 方法一： 哈希表 O(m+n)
	m1 := make(map[int]int, len(nums1))
	for _, v := range nums1 {
		m1[v] += 1
	}

	res := make([]int, 0)
	for _, v2 := range nums2 {
		if m1[v2] > 0 {
			res = append(res, v2)
			m1[v2] -= 1
		}
	}
	return res
}

// @lc code=start
func intersect(nums1 []int, nums2 []int) []int {
	// 方法 2： 排序 + 快慢指针
	sort.Ints(nums1)
	sort.Ints(nums2)

	fast, slow := 0, 0
	res := make([]int, 0)
	for fast < len(nums1) && slow < len(nums2) {
		if nums1[fast] > nums2[slow] {
			slow++
		} else if nums1[fast] < nums2[slow] {
			fast++
		} else {
			res = append(res, nums2[slow])
			fast++
			slow++
		}

	}

	return res
}
// @lc code=end

