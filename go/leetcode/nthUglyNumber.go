package main

import (
	"fmt"
)

func nthUglyNumber(n int, a int, b int, c int) int {
	indexa := 1
	indexb := 1
	indexc := 1
	result := 0
	vala := a
	valb := b
	valc := c
	for i:=0; i<n; i++ {
		a = indexa * vala
		b = indexb * valb
		c = indexc * valc
		if a <= b && a <= c {
			result = a
			indexa++
			if a == b {
				indexb++
			}
			if a == c {
				indexc++
			}
		} else if b <= a && b <= c {
			result = b
			indexb++
			if b == a {
				indexa++
			}
			if b == c {
				indexc++
			}
		} else if c <= a && c <= b {
			result = c
			indexc++
			if c == a {
				indexa++
			}
			if c == b {
				indexb++
			}
		}
	}
	return result
}
func main() {
	fmt.Println(nthUglyNumber(3, 2, 3, 5))
	fmt.Println(nthUglyNumber(1000000000, 2, 3, 4))
}