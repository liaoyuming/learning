package main

import "fmt"

func findKthNumber(n int, k int) int {
	prefix := 1
	p := 1
	for p < k {
		count := getCount(prefix, n)
		if p + count <= k {
			prefix += 1
			p += count
		} else {
			prefix *= 10
			p++
		}
	}
	return prefix

}

func getCount(prefix, max int) int {
	cur := prefix
	next := prefix + 1
	count := 0
	for cur <= max {
		count += getMin(next, max + 1) - cur
		cur *= 10
		next *= 10
	}
	return count
}

func getMin(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	fmt.Println(findKthNumber(10, 3))
}