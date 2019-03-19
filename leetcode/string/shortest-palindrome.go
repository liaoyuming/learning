package main

func shortestPalindrome(s string) string {
	if len(s) <= 0 {
		return s
	}
	m := 0
	for i:=len(s)-1; i>=0; i-- {
		if isPalindrome(&s, 0, i) {
			m = i
			break
		}
	}
	res := []byte{}
	for i:=len(s)-1; i>m; i-- {
		res = append(res, s[i])
	}
	res = append(res, s...)
	return string(res)
}

func isPalindrome(s *string, left int, right int) bool {
	for left <= right {
		if (*s)[left] != (*s)[right] {
			return false
		}
		left++
		right--
	}
	return true
}

func main() {
	println(shortestPalindrome("aacecaaadd"))
}