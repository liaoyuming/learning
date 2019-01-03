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
	fmt.Println(bubbleSort(preSort))
	//fmt.Println((time.Now().UnixNano() - pre)/1000000)
}

func bubbleSort(arr []int) []int {
	for i:=len(arr)-1; i>=0; i-- {
		for j:=len(arr)-1; j>0; j-- {
			if arr[j] < arr[j-1] {
				temp := arr[j]
				arr[j] = arr[j-1]
				arr[j-1] = temp
			}
		}
	}
	return arr
}