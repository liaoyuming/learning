package string

// KMP 算法
// 字符串查找，给定一个 haystack 字符串和一个 needle 字符串，在 haystack 字符串中找出 needle 字符串出现的第一个位置 (从0开始)。如果不存在，则返回  -1。
func kmp(haystack string, needle string) int {
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