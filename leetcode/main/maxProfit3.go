package leetcode
/**
		买卖股票的最佳时机 III
		https://leetcode-cn.com/submissions/detail/10460321/
 */
func main() {
	println(maxProfit([]int{7,1,5,3,6,4}))
}

func maxProfit(prices []int) int{
	len := len(prices)
	if len <= 0 {
		return 0
	}
	buy := [3][]int{}
	for k,_:=range buy {
		buy[k] = make([]int , len)
		buy[k][0] = -prices[0]
	}
	sell := [3][]int{}
	for k,_:=range sell {
		sell[k] = make([]int , len)
	}
	for i := 1; i < len; i++ {
		for k:=1; k < 3; k++ {
			buy[k][i] = max(sell[k-1][i-1] - prices[i], buy[k][i-1])
			sell[k][i] = max(buy[k][i-1] + prices[i], sell[k][i-1])
		}
	}
	return sell[2][len-1]
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}