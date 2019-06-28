package main

import "fmt"

func merge(intervals [][]int) [][]int {
	if len(intervals) < 2 {
		return intervals
	}
	res := [][]int{}
	intervals = sort(intervals)
	res = append(res, intervals[0])
	for i:=1; i<len(intervals); i++ {
		end := len(res)-1
		if res[end][1] < intervals[i][0] {
			res = append(res, intervals[i])
		} else if res[end][1] <= intervals[i][1] {
			res[end][1] = intervals[i][1]
		}
	}
	return res
}

func sort(arr [][]int) [][]int {
	if len(arr) < 2 {
		return arr
	}
	a := sort(arr[:len(arr)/2])
	b := sort(arr[len(arr)/2:])
	indexA := 0
	indexB := 0
	res := [][]int{}
	for indexA < len(a) || indexB < len(b) {
		if indexA >= len(a) {
			res = append(res, b[indexB])
			indexB++
		} else if indexB >= len(b) {
			res = append(res, a[indexA])
			indexA++
		} else {
			if a[indexA][0] < b[indexB][0] {
				res = append(res, a[indexA])
				indexA++
			} else {
				res = append(res, b[indexB])
				indexB++
			}
		}
	}
	return res
}

func main() {
	fmt.Print(merge([][]int{
		{1,3},
		{2,6},
		{8,10},
		{15,18},
	}))
}