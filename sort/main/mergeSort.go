package main

import (
	"math/rand"
	"fmt"
	"time"
)

func main() {
	size := 1000000
	preSort := make([]int, size)
	for i := 0; i < size; i++ {
		preSort[i] = rand.Intn(size)
	}
	//fmt.Printf("%v\n", preSort)
	pre := time.Now().UnixNano()
	mergeSort(preSort)
	//fmt.Printf("%v\n", mergeSort(preSort))

	fmt.Println((time.Now().UnixNano() - pre)/1000000)
}

func mergeSort(arr []int) []int {
	length := len(arr)
	if length < 2 {
		return arr
	}
	left := mergeSort(arr[:length/2])
	right := mergeSort(arr[length/2:])
	temp := make([]int, length)
	l := 0
	r := 0
	for l < len(left) && r < len(right) {
		if left[l] < right[r] {
			temp[l+r] = left[l]
			l++
		} else {
			temp[l+r] = right[r]
			r++
		}
	}
	for l < len(left) {
		temp[l+r] = left[l]
		l++
	}
	for r < len(right) {
		temp[l+r] = right[r]
		r++
	}
	return temp
}
