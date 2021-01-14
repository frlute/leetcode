/*
 * @lc app=leetcode id=121 lang=golang
 *
 * [121] Best Time to Buy and Sell Stock
 */

// @lc code=start
func maxProfit(prices []int) int {
	length := len(prices)
	if length < 1 {
		return 0;
	}

	// 暴力法
	// res := 0
	// for i := 0; i < length - 1; i++ {
	// 	for j := i + 1; j < length; j++ {	
	// 		profit := prices[j] - prices[i]
	// 		if profit > res {
	// 			res = profit
	// 		}		
	// 	}
	// }

	// return res

	// 锚点法
	minPrice := prices[0]
	maxProfit := 0
	for i := 1; i < length; i++ {
		if prices[i] < minPrice {
			minPrice = prices[i]
		} else if prices[i] - minPrice > maxProfit {
			maxProfit = prices[i] - minPrice
		}		
	}

	return maxProfit
}
// @lc code=end

