package main

import (
	"fmt"
)

func main() {
	var num = []int{1,1,1,2,2,3,4,4,5}
	println(removeDuplicates(&num))
	fmt.Printf("%v\n", num)


	for k := 1; k < len(num); {
		i := num[k]
		if i == num[k-1] {
			num = unset(num, k)
		} else {
			k++
		}
	}
}

func unset(a []int, i int) []int {
	copy(a[i:], a[i+1:])
	a[len(a)-1] = 0
	a = a[:len(a)-1]
	return a
}

func removeDuplicates(nums *[]int) int {
	if len(*nums) == 0 {
		return 0
	}
	i := 0;
	for j := 1; j < len(*nums); j++ {
		fmt.Printf("%v\n", *nums)
		if (*nums)[j] != (*nums)[i] {
			i++
			(*nums)[i] = (*nums)[j]
		}
	}
	return i + 1
}