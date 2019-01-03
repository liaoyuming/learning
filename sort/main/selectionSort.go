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
	fmt.Printf("%v\n", preSort)
	//pre := time.Now().UnixNano()
	fmt.Println(selectionSort(preSort))
	//fmt.Println((time.Now().UnixNano() - pre)/1000000)
}

func selectionSort(arr []int) []int {
	for i:=0; i<len(arr); i++ {
		min:=i
		for j:=i; j<len(arr); j++ {
			if arr[j] < arr[min] {
				min = j
			}
		}
		temp := arr[min]
		arr[min] = arr[i]
		arr[i] = temp
	}
	return arr
}