/*
 * @lc app=leetcode id=16 lang=golang
 *
 * [16] 3Sum Closest
 */

// @lc code=start
func threeSumClosest(nums []int, target int) int {
	length := len(nums)
	if length < 3 {
		return 0;
	}

	sort.Ints(nums);
	closest := math.MaxInt32

	abs := func(x int) int {
		if x < 0 {
			return -1 * x
		}
		return x
	}

	update := func(sum int) {
		if abs(sum - target) < abs(closest - target) {
			closest = sum 
		}
	}

	for i := 0; i < length; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		} 
		left := i + 1
		right := length - 1 
		for left < right {
			sum := nums[i]+nums[left]+nums[right]
			if sum == target {
				return sum
			} 
			update(sum)
			if sum > target {
				for left < right && nums[right] == nums[right -1] {
					right--
				}
				right--
			} else {
				for left < right && nums[left] == nums[left + 1] {
					left++
				}
				left++
			}
		}
	}

	return closest
}
// @lc code=end

