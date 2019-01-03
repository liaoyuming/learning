package leetcode

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