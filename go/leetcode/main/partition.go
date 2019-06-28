package main

import "fmt"

/**
		分割回文串
		https://leetcode-cn.com/problems/palindrome-partitioning/description/
		@todo 复习
 */

func main() {
	fmt.Printf("%v", partition("ababbbabbaba"))
}

func partition(s string) [][]string {
	len := len(s)
	// 动态规划
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
	res := [][]string{}
	dfs(s, &isPalindrome, 0, &res, []string{})
	return res
}
// 回溯， 深度搜索
func dfs(s string, isPalindrome *[][]bool, start int, res *[][]string, tempRes []string) {
	if start >= len(s) {
		*res = append(*res, tempRes)
		return
	}
	for i:=start; i<len(s); i++ {
		if (*isPalindrome)[start][i] {
			tempRes = tempRes[:len(tempRes):len(tempRes)] // golang  这里有坑
			dfs(s, isPalindrome, i+1, res, append(tempRes, s[start:i+1]))
		}
	}
}
