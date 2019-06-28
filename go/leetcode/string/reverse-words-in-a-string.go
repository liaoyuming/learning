package main

func reverseWords(s string) string {
	if len(s) < 1 {
		return s
	}
	res := ""
	start := 0
	for i:=0; i<len(s); i++ {
		if s[i] == ' ' {
			if i > start {
				res = s[start:i] + " " + res
			}
			start = i+1
			continue
		}
	}
	if start < len(s) {
		res = s[start:] + " " + res
	}
	if len(res) > 0 {
		res = res[:len(res)-1]
	}
	return res
}