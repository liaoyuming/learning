package main

func maxArea(height []int) int {
	if len(height) < 2 {
		return 0
	}
	max := 0
	for i,j:=0,len(height)-1; i < len(height) && j>i; {
		temp := 0
		if height[i] < height[j] {
			temp = (j-i)*height[i]
			i++
		} else {
			temp = (j-i)*height[j]
			j--
		}
		if max < temp {
			max = temp
		}
	}
	return max
}
