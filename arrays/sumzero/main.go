package main

import "fmt"

func subarraySum(nums []int, k int) int {
	cache := make(map[int]int)

	res := 0
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		if sum == k {
			res += 1
		}

		if v, ok := cache[sum-k]; ok {
			res += v
		}

		if v, ok := cache[sum]; !ok {
			cache[sum] = 1
		} else {
			cache[sum] = v + 1
		}
	}
	fmt.Println(cache)
	return res
}

func main() {
	fmt.Println(subarraySum([]int{-3, 3, 2, 3}, 0))
}
