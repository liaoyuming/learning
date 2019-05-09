package main

func nthUglyNumber(n int) int {
	dp := make([]int, n+1)
	dp[1] = 1
	tmp2 := 1
	tmp3 := 1
	tmp5 := 1
	for i:=2; i<=n; i++ {
		a := dp[tmp2]*2
		b := dp[tmp3]*3
		c := dp[tmp5]*5
		minVal := min(a, b, c)
		if a == minVal {
			tmp2++
		}
		if b == minVal {
			tmp3++
		}
		if c == minVal {
			tmp5++
		}
		dp[i] = minVal
	}
	return dp[n]
}

func min(a, b, c int) int {
	if a <= b && a <= c {
		return a
	}
	if b <= a && b <= c {
		return b
	}
	return c
}

func main() {
	println(nthUglyNumber(10))
}