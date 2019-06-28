package main

import "math"

/**
		整数反转
		https://leetcode-cn.com/problems/reverse-integer/description/
 */

func main() {
	println(reverse(1534236469))
}

func reverse(x int) int {
	res := 0
	for x != 0 {
		res += x%10
		res*=10
		x=x/10
	}
	res/=10
	res += x
	if res < -int(math.Pow(2, 31)) || res > int(math.Pow(2, 31) - 1) {
		return 0
	}
	return res
}