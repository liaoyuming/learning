package main

func getSum(a int, b int) int {
	m := a ^ b
	n := (a & b) << 1
	if n == 0 {
		return m
	}
	return getSum(m, n)
}