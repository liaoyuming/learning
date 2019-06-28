package main

/**
		买卖股票的最佳时机 IV
		https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-iv/description/
 */

func main() {
	println(maxProfit(2, []int{3,3,5,0,0,3,1,4}))
	//println(maxProfit(1000000000, []int{106,373,495,46}))
}

func maxProfit3(k int, prices []int) int {
	len := len(prices)
	if len <= 0 {
		return 0
	}
	//当交易次数大于天数，看作是贪心算法
	if k > len {
		res := 0
		for i := 1 ; i < len; i++ {
			if prices[i] > prices[i-1] {
				res += prices[i]-prices[i-1]
			}
		}
		return res;
	}
	buy := make([][]int, k+1)
	for key,_:=range buy {
		buy[key] = make([]int , len)
		buy[key][0] = -prices[0]
	}
	sell := make([][]int, k+1)
	for key,_:=range sell {
		sell[key] = make([]int , len)
	}
	for i := 1; i < len; i++ {
		for t:=1; t < k+1; t++ {
			buy[t][i] = max(sell[t-1][i-1] - prices[i], buy[t][i-1])
			sell[t][i] = max(buy[t][i-1] + prices[i], sell[t][i-1])
		}
	}
	return sell[k][len-1]
}

func maxProfit(k int, prices []int) int {
	len := len(prices)
	if len <= 0 {
		return 0
	}
	if k > len {
		res := 0
		for i := 1 ; i < len; i++ {
			if prices[i] > prices[i-1] {
				res += prices[i]-prices[i-1]
			}
		}
		return res;
	}

	buy := make([]int, k+1)
	sell := make([]int, k+1)
	for key,_:=range buy {
		buy[key] = -prices[0]
	}
	for i := 1; i < len; i++ {
		for t:=1; t < k+1; t++ {
			buy[t] = max(sell[t-1] - prices[i], buy[t])
			sell[t] = max(buy[t] + prices[i], sell[t])
		}
	}
	return sell[k]
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}