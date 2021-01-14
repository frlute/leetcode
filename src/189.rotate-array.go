/*
 * @lc app=leetcode id=189 lang=golang
 *
 * [189] Rotate Array
 */

// @lc code=start
func rotate(nums []int, k int)  {
	length := len(nums)
	if length < 2 {
		return
	}
	// 考虑极限场景，k > length

	// 方法一： 暴力法
	// for i := 0; i < k; i++ {
	// 	for j:= 0; j< length; j++ {
	// 		nums[j], nums[length-1] = nums[length-1], nums[j] 
	// 	}
	// }

	// 方法二： 使用额外的数组, 也就是需要把 i 位置的元素移动到 i+k/len(nums)
	// tmp := make([]int, length)
	// for i, value := range nums {
	// 	tmp[(i+k)%length] = value
	// }
	
	// for j := range nums  {
	// 	nums[j] = tmp[j]
	// }

	// 方法三： 使用环状替换


	// 方法四： 使用反转
	// reverse := func (list []int)  {
	// 	length := len(list)
	// 	for i:=0; i < length/2; i++ {
	// 		list[i], list[length-i-1] = list[length-i-1], list[i]
	// 	}
	// }
	
	// reverse(nums)
	// reverse(nums[:k%length])
	// reverse(nums[k%length:])
}
// @lc code=end

