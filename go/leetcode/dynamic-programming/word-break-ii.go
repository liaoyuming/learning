package main

import "strings"

func wordBreak(s string, wordDict []string) []string {
	length := len(s)
	dp := make([][]string, length+1)
	wd := make(map[string]string, len(wordDict))
	for _, v := range wordDict {
		wd[v] = v
	}
	for i:=1; i<length+1; i++ {
		for j:=i-1; j>=0; j-- {
			if _, ok := wd[s[j:i]]; ok && (j==0 || dp[j] != nil){
				if j == 0 {
					dp[i] = append(dp[i], s[j:i])
				} else {
					t := dp[i]
					for _, v := range dp[j] {
						t = append(t, v + " " + s[j:i])
					}
					dp[i] = t
				}
			}
		}
	}
	return dp[len(s)]
}

func wordBreak2(s string, wordDict []string) []string {
	return dfs(s, wordDict, map[string][]string{})
}

func dfs(s string, wordDict []string, tmp map[string][]string) []string {
	if _, ok := tmp[s]; ok {
		return tmp[s]
	}
	if len(s) == 0 {
		return []string{""}
	}
	res := []string{}
	for _, word := range wordDict {
		if strings.HasPrefix(s, word) {
			sublist := dfs(s[len(word):], wordDict, tmp)
			for _, sub := range sublist {
				if sub != ""{
					res = append(res, word + " " + sub)
				} else {
					res = append(res, word)
				}
			}
		}
	}
	tmp[s] = res
	return res
}

func main() {
	//words := wordBreak3("catsanddog", []string{"cat", "cats", "and", "sand", "dog"})
	words := wordBreak("aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa", []string{"a","aa","aaa","aaaa","aaaaa","aaaaaa","aaaaaaa","aaaaaaaa","aaaaaaaaa","aaaaaaaaaa"})
	for _, w:=range words{
		println(w)
	}
}