/*
 * @lc app=leetcode id=26 lang=golang
 *
 * [26] Remove Duplicates from Sorted Array
 */

 func removeDuplicates1(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}
	for i := 0; i+1 < len(nums); {
		if nums[i] == nums[i+1] {
			nums = append(nums[:i], nums[i+1:]...)
		} else {
			i++
		}
	}

	return len(nums)
}


// @lc code=start
func removeDuplicates(nums []int) int {
	// 双指针法，也称为快慢指针
	if len(nums) <= 1 {
		return len(nums)
	}
	i := 0
	for j := 1; j < len(nums); j++ {
		// 此时数组为 [1, 2, 3] 等完全不重复的数组时，
		// 还是会产生一次复制
		// 可以增加 j - i > 1 时再赋值 
		/*
		if j - i > 1 {
			nums[i] = nums[j]
		}
		*/
		if nums[j] != nums[i] {
			i++
			nums[i] = nums[j]
		}
	}
	fmt.Println(nums)

	return i + 1
}
// @lc code=end
