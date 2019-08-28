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
	heapSort(preSort)
	//fmt.Printf("%v\n", heapSort(preSort))
	//
	fmt.Println((time.Now().UnixNano() - pre)/1000000)
}


func heapSort(arr []int) []int {
	for i:=len(arr)/2-1; i>=0; i-- {
		arr = buildHeap(arr, i, len(arr))
	}
	for i:=len(arr)-1; i>0; i-- {
		arr[0] , arr[i] = arr[i], arr[0]
		arr = buildHeap(arr, 0, i)
	}
	return arr
}

func buildHeap(arr []int, i int, length int) []int {
	left := 2*i+1
	right := 2*i+2
	large := i
	if left < length && arr[left] > arr[large] {
		large = left
	}
	if right < length && arr[right] > arr[large] {
		large = right
	}
	if large != i {
		arr[i], arr[large] = arr[large], arr[i]
		return buildHeap(arr, large, length)
	}
	return arr
}