package main

import "strconv"

func maximumSwap(num int) int {
	strNum := []byte(strconv.Itoa(num))
	maxIndex := -1
	for i:=1; i<len(strNum); i++ {
		if strNum[i] > strNum[i-1] && (maxIndex == -1 || strNum[maxIndex] < strNum[i]) {
			maxIndex = i
		}
	}
	if maxIndex == -1 {
		return num
	}

	for i:=0; i<len(strNum); i++ {
		if strNum[i] < strNum[maxIndex] {
			strNum[i], strNum[maxIndex] = strNum[maxIndex], strNum[i]
			break
		}
	}
	res, _ := strconv.Atoi(string(strNum))
	return res
}