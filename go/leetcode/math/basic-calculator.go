package main

import (
	"strconv"
	"fmt"
)

func calculate(s string) int {
	tmp := []byte{}
	for i:=0; i<len(s); i++ {
		if s[i] == ' ' {
			continue
		}
		if s[i] == ')' {
			t := ""
			r := 0
			for len(tmp) > 0 {
				cur := tmp[len(tmp)-1]
				tmp = tmp[:len(tmp)-1]
				if cur == '(' {
					r = calculate(t)
					break
				}
				t = string(cur) + t
			}
			tmp = append(tmp, strconv.Itoa(r)...)
		} else {
			tmp = append(tmp, s[i])
		}
		fmt.Println(string(tmp))
	}
	a := ""
	b := ""
	flag := false
	var op byte = 0
	for i:=0; i<len(tmp); i++ {
		if tmp[i] == '-' || tmp[i] == '+' {
			if i-1 > 0 {
				if (tmp[i-1] == '-' && tmp[i] == '-') || (tmp[i-1] == '+' && tmp[i] == '+'){
					op = '+'
					tmp[i] = op
					continue
				} else if (tmp[i-1] == '+' && tmp[i] == '-') || (tmp[i-1] == '-' && tmp[i] == '+') {
					op = '-'
					tmp[i] = op
					continue
				}
			}
			if flag {
				a = cal(a, b, op)
				b = ""
			} else {
				flag = true
			}
			op = tmp[i]
		} else {
			if !flag {
				a = a + string(tmp[i])
			} else {
				b = b + string(tmp[i])
			}
		}
	}
	if op != 0 {
		a = cal(a, b , op)
	}
	res, _ := strconv.Atoi(a)
	return res
}

func cal(a string, b string, op byte) string {
	ia, _ := strconv.Atoi(a)
	ib, _ := strconv.Atoi(b)
	res := ""
	if op == '-' {
		res = strconv.Itoa(ia - ib)
	} else if op == '+' {
		res = strconv.Itoa(ia + ib)
	}
	return res
}


func calculate2(s string) int {
	res := 0
	stack := make([]int, 0, len(s))
	sign := 1
	num := 0

	for i := 0; i < len(s); i++ {
		switch s[i] {
		case '1', '2', '3', '4', '5', '6', '7', '8', '9', '0':
			// 提取 s 中的数字
			num = 0
			for ; i < len(s) && s[i] >= '0' && s[i] <= '9'; i++ {
				num = 10*num + int(s[i]-'0')
			}
			// 根据前面的记录，进行运算
			res += sign * num
			// 此时 s[i] 已经不是数字了
			// for 语句中，会再＋1，所以这里先 -1
			i--
		case '+', '-':
			if i-1 > 0 && (s[i-1] == '-' || s[i-1] == '+') {
				if s[i] == '-' {
					sign = -1 * sign
				}
			} else {
				if s[i] == '+' {
					sign = 1
				} else {
					sign = -1
				}
			}
		case '(':
			stack = append(stack, res, sign)
			res = 0
			sign = 1
		case ')':
			sign = stack[len(stack)-1]
			temp := stack[len(stack)-2]
			stack = stack[:len(stack)-2]
			res = sign*res + temp
		}
	}

	return res
}

func main() {
	println(calculate2(" 2-1 +-+- 2 "))
}