package main

import (
	"strconv"
)

/**
		回文数
		https://leetcode-cn.com/problems/palindrome-number/
 */

func main() {
	println(isPalindrome2(11))
}

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	z := x
	y := 0
	for x != 0 {
		y += x%10
		y *= 10
		x = x/10
	}
	y /= 10
	return y == z
}

func isPalindrome2(x int) bool {
	if x < 0 {
		return false
	}
	s := strconv.Itoa(x)
	for i:=0; i<=len(s)/2; i++ {
		if s[i] != s[len(s)-i-1] {
			return false
		}
	}
	return true
}

