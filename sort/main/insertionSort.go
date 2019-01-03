package main

import (
	"math/rand"
	"fmt"
)

func main() {
	size := 10
	preSort := make([]int, size)
	for i := 0; i < size; i++ {
		preSort[i] = rand.Intn(size*10)
	}
	fmt.Println(preSort)
	//pre := time.Now().UnixNano()
	fmt.Println(insertionSort(preSort))
	//fmt.Println((time.Now().UnixNano() - pre)/1000000)
}

func insertionSort(arr []int) []int {
	for i:=1; i<len(arr); i++ {
		preIndex := i-1
		current := arr[i]
		for preIndex>=0 && arr[preIndex] > current {
			arr[preIndex+1] = arr[preIndex]
			preIndex--
		}
		arr[preIndex+1] = current
	}
	return arr
}