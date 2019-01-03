package leetcode

// 有效的括号

func main() {
	res := isValid("")
	println(res)
}

func isValid(s string) bool {
	l := len(s)
	n := 0
	temp := make([]uint8, len(s))
	for i:=0; i<l; i++{
		if s[i] == "("[0] || s[i] == "["[0] || s[i] == "{"[0] {
			temp[n] = s[i]
			n++
			continue
		}
		if n < 1 {
			return false
		}
		if (s[i] == ")"[0] && temp[n-1] == "("[0]) || (s[i] == "]"[0] && temp[n-1] == "["[0]) || (s[i] == "}"[0] && temp[n-1] == "{"[0]) {
			n--
		} else {
			return false
		}
	}
	return n == 0
}