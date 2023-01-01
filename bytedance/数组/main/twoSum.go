package main

import "fmt"

func main() {
	nums := []int{3, 2, 4}
	target := 6

	fmt.Println(twoSum(nums, target))
}

func twoSum(nums []int, target int) []int {
	mp := make(map[int]int)
	for k, v := range nums {
		if mp[target-v] != 0 {
			return []int{mp[target-v] - 1, k}
		}
		mp[v] = k + 1
	}
	return []int{}
}
