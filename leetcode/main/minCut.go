package main

import "fmt"

/**
		分割回文串 II
		https://leetcode-cn.com/problems/palindrome-partitioning-ii/description/
 */

func main() {

	fmt.Println(minCut("aabcbcc"))
}


func minCut(s string) int {
	len := len(s)
	isPalindrome := make([][]bool, len)
	for k,_:=range isPalindrome {
		isPalindrome[k] = make([]bool, len)
		isPalindrome[k][k] = true
		if k < len-1 {
			isPalindrome[k][k+1] = s[k] == s[k+1]
		}
	}
	for i:=len-3; i>=0; i-- {
		for j:=i+2; j<len; j++ {
			isPalindrome[i][j] = isPalindrome[i+1][j-1] && s[i] == s[j]
		}
	}
	cut := make([]int, len+1)
	cut[0] = 0
	for i:=1; i<=len; i++ {
		cut[i] = i
		for j:=i-1; j>=0; j-- {
			if isPalindrome[j][i-1] {
				cut[i] = min(cut[i], cut[j]+1)
			}
		}
	}
	return cut[len]-1
}

func min(a int, b int) int {
	if a<=b {
		return a
	}
	return b
}