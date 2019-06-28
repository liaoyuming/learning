package main

func main() {
	println(maxProfit([]int{7,1,5,3,6,4}))
}

/**
	https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock/description/
 */
func maxProfit(prices []int) int {
	if len(prices) < 2 {
		return 0
	}
	max := 0
	for key, buy := range prices {
		if key+1 > len(prices) {
			return max
		}
		temp := prices[key+1:]
		for _,  sale:= range temp {
			earn := sale - buy
			if earn > max {
				max = earn
			}
		}
	}
	return max
}

func maxProfit(prices []int) int {
	if len(prices) < 2 {
		return 0
	}
	dp := make([]int, len(prices))
	dp[0] = prices[0]
	max := 0
	for i:=1; i<len(prices); i++ {
		if prices[i] < dp[i-1] {
			dp[i] = prices[i]
		} else {
			dp[i] = dp[i-1]
		}
		t := prices[i] - dp[i-1]
		if max < t {
			max = t
		}
	}
	return max
}

func maxProfit(prices []int) int {
	if len(prices) < 2 {
		return 0
	}
	lowest := prices[0]
	max := 0
	for i:=1; i<len(prices); i++ {
		t := prices[i] - lowest
		if max < t {
			max = t
		}
		if prices[i] < lowest {
			lowest = prices[i]
		}
	}
	return max
}

