package main

/**
		罗马数字转整数
		https://leetcode-cn.com/problems/roman-to-integer/
 */

func main() {
	println(romanToInt("MCMXCIV"))
}

func romanToInt(s string) int {
	config := map[byte]int{
		'I':1,
		'V':5,
		'X':10,
		'L':50,
		'C':100,
		'D':500,
		'M':1000,
	}
	len := len(s)
	num := 0
	for i:=0; i<len-1; i++ {
		if config[s[i]] < config[s[i+1]] {
			num -= config[s[i]]
		} else {
			num += config[s[i]]
		}
	}
	num += config[s[len-1]]
	return num
}