/*
 * @lc app=leetcode id=27 lang=golang
 *
 * [27] Remove Element
 */

func removeElement1(nums []int, val int) int {
	numsLen := len(nums)
	if (numsLen <= 1) {
		return numsLen
	}

	i := 0
	for j := 0; j < numsLen; j++ {
		if nums[j] != val {
			nums[i]=nums[j]
			i++
		}
	} 
	
	return i
}

/*
上面方法如遇到 [1，2，3，5，4] 删除 4 时，会产生不必要 copy 的问题
如遇到 [4, 1, 2, 3, 5] 删除 4 时似乎没必要左移 [1, 2, 3, 5]

当我们遇到 nums[i] = valnums[i]=val 时，我们可以将当前元素与最后一个元素进行交换，并释放最后一个元素。这实际上使数组的大小减少了 1。
请注意，被交换的最后一个元素可能是您想要移除的值。但是不要担心，在下一次迭代中，我们仍然会检查这个元素。
*/
// @lc code=start
func removeElement(nums []int, val int) int {
	for i := 0; i < len(nums); {
		if nums[i] == val {
			nums = append(nums[:i], nums[i+1:]...)
		} else {
			i++
		}
	}

	return len(nums)
}
// @lc code=end