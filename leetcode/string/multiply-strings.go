package main

import (
	"time"
)

func multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}
	len1, len2 := len(num1), len(num2)
	t := make([]int, len1 + len2)
	for i := len2 - 1; i >= 0; i-- {
		for j := len1 - 1; j >= 0; j-- {
			t[len1+len2-2-i-j] += int(num1[j]-"0"[0]) * int(num2[i]-"0"[0])
		}
	}
	res := ""
	for i:=0; i<len(t)-1; i++ {
		t[i+1] += t[i] / 10
		t[i] %= 10
		res = string(byte(t[i]) + "0"[0]) + res
	}
	if t[len(t)-1] != 0 {
		res = string(byte(t[len(t)-1]) + "0"[0]) + res
	}
	return res
}

func main() {
	println(time.Now().String())
	println(multiply("100000000000000000","10000000000000000000000000000"))
	println(time.Now().String())

}