package main

import (
	"math/rand"
	"time"
	"fmt"
)

func main() {
	size := 1000000
	preSort := make([]int, size)
	for i := 0; i < size; i++ {
		preSort[i] = rand.Intn(size)
	}
	//fmt.Printf("%v\n", preSort)
	pre := time.Now().UnixNano()
	//quickSort(preSort)
	quickSort2(&preSort, 0, len(preSort)-1)
	//fmt.Printf("%v\n", preSort)
	fmt.Println((time.Now().UnixNano() - pre)/1000000)
}

func quickSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	k := arr[0]
	var a, b []int
	for i:=1; i<len(arr); i++ {
		if arr[i] < k {
			a = append(a, arr[i])
		} else {
			b = append(b, arr[i])
		}
	}
	return append(append(quickSort(a), k), quickSort(b)...)
}

func quickSort2(arr *[]int, start int, end int) {
	if start >= end {
		return
	}
	x := (*arr)[start]
	a := start
	b := end
	for a < b {
		for a < b && (*arr)[b] >= x {
			b--
		}
		if a < b {
			(*arr)[a] = (*arr)[b]
		}
		for a < b && (*arr)[a] < x {
			a++
		}
		if a < b {
			(*arr)[b] = (*arr)[a]
		}
	}
	(*arr)[a] = x
	quickSort2(arr, start, a-1)
	quickSort2(arr, a+1, end)
}