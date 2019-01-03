package leetcode
/**
   	url: https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-with-cooldown/description/
 */
func main() {
	println(maxProfit3([]int{1,2,3,0,2}))
}

func maxProfit(prices []int) int {
	max := 0
	return maxProfitIter(prices, 0, 1, max)
}

func maxProfitIter(prices []int, start int, buy int, res int) int{
	if start > len(prices)-2 || buy > len(prices)-1 {
		return res
	}
	sale := prices[buy] - prices[start]
	temp1 := maxProfitIter(prices, buy+2, buy+3, res + sale)
	temp2 := maxProfitIter(prices, start, buy+1, res)
	temp3 := maxProfitIter(prices, start+1, start+2, res)
	return max(max(temp1, temp2), temp3)
}

func maxProfit2(prices []int) int{
	len := len(prices)
	if len <=0 {
		return 0
	}
	buy := make([]int , len+1)
	sell := make([]int , len+1)
	rest := make([]int , len+1)
	buy[0] = -prices[0]
	for i := 0; i < len; i++ {
		buy[i+1] = max(rest[i] - prices[i], buy[i])
		sell[i+1] = max(buy[i] + prices[i], sell[i])
		rest[i+1] = max(max(sell[i], buy[i]), rest[i])
	}
	return sell[len]
}

// 由rest[i] = sell[i-1]精简为两个状态
func maxProfit3(prices []int) int {
	len := len(prices)
	if len <=0 {
		return 0
	}
	buy := make([]int , len+1)
	sell := make([]int , len+1)
	buy[1] = -prices[0]
	for i := 1; i < len; i++ {
		buy[i+1] = max(sell[i-1] - prices[i], buy[i])
		sell[i+1] = max(buy[i] + prices[i], sell[i])
	}
	return sell[len]
}



func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
