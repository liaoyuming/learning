package main

func threeSumClosest(nums []int, target int) int {
	if len(nums) < 3 {
		sum := 0
		for _,v:=range nums {
			sum += v
		}
		return sum
	}
	nums = quickSort(nums)
	res := nums[0] + nums[1] + nums[2]
	for i := 1; i < len(nums)-1; i++ {
		for j,k :=0,len(nums)-1; j<i && i<k; {
			t := nums[j] + nums[k] + nums[i]
			if  t > target {
				k--
			} else if t == target {
				return t
			} else {
				j++
			}
			if abs(target - t) < abs(target - res) {
				res = t
			}
		}
	}
	return res
}

func abs(res int) int {
	if res < 0 {
		return - res
	}
	return res
}

func quickSort(nums []int) []int {
	if len(nums) < 2 {
		return nums
	}
	k := nums[0]
	a := []int{}
	b := []int{}
	for i:=1; i<len(nums); i++ {
		if nums[i] < k {
			a = append(a, nums[i])
		} else {
			b = append(b, nums[i])
		}
	}
	return append(append(quickSort(a), k), quickSort(b)...)
}

func main() {
	println(threeSumClosest([]int{-1,2,1,-4}, 1))
}