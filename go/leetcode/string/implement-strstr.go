package main

// 暴力求解
func strStr2(haystack string, needle string) int {
	if len(needle) < 1 {
		return 0
	}
	if len(haystack) < 1 {
		return -1
	}
	for i:=0; i<=len(haystack)-len(needle); i++ {
		for j:=0; j<len(needle); j++ {
			if haystack[i+j] != needle[j] {
				break
			}
			if j == len(needle)-1 {
				return i
			}
		}
	}
	return -1
}

// KMP 算法
func strStr(haystack string, needle string) int {
	if len(needle) < 1 {
		return 0
	}
	if len(haystack) < 1 {
		return -1
	}
	next := getNext(needle)
	i :=0
	j := 0
	for i<len(haystack) && j<len(needle) {
		if j == -1 || haystack[i] == needle[j] {
			i++
			j++
		} else {
			j = next[j]
		}
	}
	if j == len(needle) {
		return i-j
	}
	return -1
}

func getNext(needle string) []int {
	next := make([]int, len(needle))
	next[0] = -1
 	k := -1
	j := 0
	for j < len(needle)-1 {
		if k == -1 || needle[k] == needle[j] {
			k++
			j++
			next[j] = k
			if needle[j] == needle[k] {
				next[j] = next[k]
			}
		} else {
			k = next[k]
		}
	}
	return next
}



func main() {
	println(strStr("aabaa", "baa"))
}