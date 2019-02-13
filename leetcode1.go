package main

import "fmt"

func main() {
	nums := []int{2, 7, 11, 15}
	target := 18
	indexs := twoSum(nums, target)
	fmt.Printf("indexs: %d, %d", indexs[0], indexs[1])
}

func twoSum(nums []int, target int) []int {
	numMap := make(map[int]int)
	numLen := len(nums)
	for i := 0; i < numLen; i++ {
		if index, ok := numMap[target-nums[i]]; ok {
			return []int{index, i}
		} else {
			numMap[nums[i]] = i
		}
	}
	return []int{}
}
