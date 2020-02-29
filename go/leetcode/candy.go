package main

import (
	"fmt"
)

func candy(ratings []int) int {
	if len(ratings) == 0 {
		return 0
	} else if len(ratings) == 1 {
		return 1
	}
	count :=0
	cache := make([]int, len(ratings))
	for i:=0; i<len(ratings); i++{
		count += getCurrentCandy(ratings, i, &cache)
	}
	return count
}

func getCurrentCandy(ratings []int, i int, cache *[]int) int {
	if (*cache)[i] > 0 {
		return (*cache)[i]
	}
	res := 1
	if i == 0 {
		if ratings[i] > ratings[i+1] {
			res = getCurrentCandy(ratings, i+1, cache) + 1
		}
	} else if i == len(ratings)-1 {
		if ratings[i] > ratings[i-1] {
			res = getCurrentCandy(ratings, i-1, cache) + 1
		}
	} else if ratings[i] > ratings[i-1] && ratings[i] > ratings[i+1] {
		res = getMax(getCurrentCandy(ratings, i-1, cache), getCurrentCandy(ratings, i+1, cache)) + 1
	} else if ratings[i] > ratings[i-1] {
		res = getCurrentCandy(ratings, i-1, cache)+1
	} else if ratings[i] > ratings[i+1] {
		res = getCurrentCandy(ratings, i+1, cache) +1
	}
	(*cache)[i] = res
	return res
}

func getMax(a int, b int) int {
	if a < b {
		return b
	}
	return a
}

func main()  {
	fmt.Println(candy1([]int{1, 3, 4, 5, 2}))
}


func candy1(ratings []int) int {
	candy := make([]int, len(ratings))
	candy[0] = 1
	for i:=1; i< len(ratings); i++ {
		if ratings[i] > ratings[i-1] {
			candy[i] = candy[i-1] + 1
		} else {
			candy[i] = 1
		}
	}
	count := candy[len(ratings)-1]
	for i:=len(ratings)-2; i>=0; i-- {
		if ratings[i] > ratings[i+1] {
			candy[i] = getMax(candy[i+1]+1, candy[i])
		}
		count += candy[i]
	}
	return count
}