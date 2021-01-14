/*
 * @lc app=leetcode id=18 lang=golang
 *
 * [18] 4Sum
 */

// @lc code=start
func fourSum(nums []int, target int) [][]int {
	length := len(nums)
	res := make([][]int, 0)
	if length < 4 {
		return res
	}
	sort.Ints(nums);

	for i := 0; i < length - 3; i++ {
		// if nums[i] > target {
		// 	return res
		// }
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		for j := i + 1; j < length - 2; j++ {
			if j > i + 1 && nums[j] == nums[j-1] {
				continue
			}
			left, right := j + 1, length - 1
			for left < right {
				sum := nums[i] + nums[j] + nums[left] + nums[right]
				if sum == target {
					res = append(res, []int{nums[i], nums[j], nums[left], nums[right]})
					left++
					right--
					for left < right && nums[left] == nums[left - 1] {
						left++
					}
					for left < right && nums[right] == nums[right + 1] {
						right--
					}
				} else if sum > target {
					right--
				} else {
					left++
				}
			}
		}
	}
	return res
}
// @lc code=end

