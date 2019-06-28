package main
/**
	https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-with-transaction-fee/description/
 */

func main()  {
	println(maxProfit([]int{1, 3, 2, 8, 4, 9}, 2))
}

func maxProfit(prices []int, fee int) int {
	len := len(prices)
	if len <= 0 {
		return 0
	}
	buy := make([]int, len)
	sell := make([]int, len)
	buy[0] = -prices[0]
	for i:=1; i<len; i++ {
		buy[i] = max(sell[i-1] - prices[i], buy[i-1])
		sell[i] = max(buy[i-1] + prices[i] - fee, sell[i-1])
	}
	return sell[len-1]
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}