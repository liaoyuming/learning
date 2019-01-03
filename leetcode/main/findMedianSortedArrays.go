package main

import "fmt"

func main() {
	nums1 := []int{1, 3}
	nums2 := []int{2, 4}
	res := findMedianSortedArrays(nums1, nums2)
	fmt.Println(res)
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	l1 := len(nums1)
	l2 := len(nums2)
	medianIndex := (l1 + l2) / 2
	nums1Index := 0
	nums2Index := 0
	sortedArray := make([]int, medianIndex + 1)
	for i := 0; i <= medianIndex; i++ {
		if nums1Index >= l1 {
			sortedArray[i] = nums2[nums2Index]
			nums2Index++
			continue
		}
		if nums2Index >= l2 {
			sortedArray[i] = nums1[nums1Index]
			nums1Index++
			continue
		}
		if nums1[nums1Index] <= nums2[nums2Index] {
			sortedArray[i] = nums1[nums1Index]
			nums1Index++
		} else {
			sortedArray[i] = nums2[nums2Index]
			nums2Index++
		}
	}
	if (l1 + l2) % 2 == 1 {
		return float64(sortedArray[medianIndex])
	} else {
		return float64((sortedArray[medianIndex] + sortedArray[medianIndex - 1])) / 2
	}
}