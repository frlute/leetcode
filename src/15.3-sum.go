/*
 * @lc app=leetcode id=15 lang=golang
 *
 * [15] 3Sum
 */

// @lc code=start
func threeSum(nums []int) [][]int {
	length := len(nums)
	if length < 3 {
		return [][]int{}
	}

	sort.Ints(nums);
	res := make([][]int, 0)
	expectValue := 0

	for i := 0; i < length; i++ {
		if nums[i] > expectValue {
			return res
		}
		if i > 0 && nums[i] == nums[i-1] {
			continue
		} 
		left := i + 1
		right := length - 1 
		for left < right {
			sum := nums[i]+nums[left]+nums[right]
			if sum == expectValue {
				res = append(res, []int{nums[i], nums[left], nums[right]})
				for left < right && nums[left] == nums[left+1]{
					left++
				}
				for left < right && nums[right] == nums[right-1] {
					right--
				}
				right--
				left++
			} else if sum > expectValue {
				right--
			} else {
				left++
			}
		}
	}

	return res
}
// @lc code=end

