package string

// KMP 算法， 时间复杂度 O(m+n)
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
			// 从 next 获取到右移的位置
			j = next[j]
		}
	}
	// 匹配成功
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
		// needle[k] 前缀
		// needle[j] 后缀
		if k == -1 || needle[k] == needle[j] {
			// 如果 needle 中 [0, k] 的值 和 [j-k, j] 的值依次相同，则 next[j+1] = k+1 ，此为充分必要条件
			k++
			j++
			next[j] = k
			// 如果 needle[j] 和 needle[next[j]] 相同，则取 next[next[j]] 即 next[k]
			// 原因：如果 haystack[i] != needle[j]，若 needle[j] == needle[next[j]]， 则必定 haystack[i] != needle[next[j]]，因此需要避免这种重复匹配
			if needle[j] == needle[k] {
				next[j] = next[k]
			}
		} else {
			// 因为 k < j，所以存在 next[k] = k'，即 needle 中 [0, k'-1] 的值 和 [k-k', k-1] 的值依次相同
			// 因为 next[j] = k 即 [0, k-1] 的值 和 [j-k, j-1] 的值依次相同
			// 又 0 < k' < k
			// 因此，必然有 [0, k'-1] 与 [j-k', j-1] 的值依次相同
			// 如果 needle[k'] 与 needle[j] 相同， 则 [0, k'] 的值 和 [j-k', j] 的值依次相同，也就是 next[j+1] = k'+1
			// 所以，这里可以直接用 next[k] 赋值给 k
			k = next[k]
		}
	}
	return next
}