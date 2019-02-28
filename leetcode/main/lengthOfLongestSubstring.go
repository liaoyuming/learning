package main

func main() {
	l := lengthOfLongestSubstring("pwwkew")
	println(l)
}

func lengthOfLongestSubstring(s string) int {
	max := 0
	l := len(s)
	temp := [256]int{}
	n := 0
	t := 0
	for i:=0; i < l; i++ {
		if v := temp[s[i]]; n < v {
			n = v
		}
		temp[s[i]] = i + 1
		t = temp[s[i]] - n
		if max < t {
			max = t
		}
	}
	return max
}
