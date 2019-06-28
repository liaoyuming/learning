package main

func main() {
	res := longestPalindrome("cbbd")
	println(res)
}

func longestPalindrome(s string) string {
	l := len(s)
	if l < 2 {
		return s
	}
	var (
		len1 int
		len2 int
	)
	len := 0
	start := 0
	end := 0
	for i:=0; i<l-(len-1)/2; i++ {
		len1 = getMaxLen(&s, i, i, &l)
		len2 = getMaxLen(&s, i, i+1, &l)
		if len1 < len2 {
			len = len2
		} else {
			len = len1
		}
		if len > end - start {
			start = i - (len - 1) / 2
			end = i + len/2
		}
	}
	return s[start:end+1]
}

func getMaxLen(s *string, left int, right int, l *int) int {
	for left >= 0 && right < *l &&  (*s)[left] == (*s)[right] {
		left--
		right++
	}
	return right - left - 1
}