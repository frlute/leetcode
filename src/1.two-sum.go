package main

/*
 * @lc app=leetcode id=1 lang=golang
 *
 * [1] Two Sum
 */

// 暴力法
// func twoSum(nums []int, target int) []int {
// 	numsLen := len(nums)
// 	for i, value := range nums {
// 		for j := i + 1; j < numsLen; j++ {
// 			if value+nums[j] == target {
// 				return []int{i, j}
// 			}
// 		}
// 	}
// 	return []int{}
// }

// člc code=start
func twoSum(nums []int, target int) []int {
	dist := make(map[int]int, len(nums))
	for index, value := range nums {
		expectValue := target - value
		if expectIndex, ok := dist[expectValue]; ok && expectIndex != index {
			return []int{expectIndex, index}
		}
		dist[value] = index
	}

	return nil
}
