package main

import "fmt"

// Given an array of integers, return indices of the two numbers such that they add up to a specific target.

// You may assume that each input would have exactly one solution, and you may not use the same element twice.

// Example:

// Given nums = [2, 7, 11, 15], target = 9,

// Because nums[0] + nums[1] = 2 + 7 = 9,
// return [0, 1].

// 要从数组中找到两个数据，使得两数之和等于目标值，输出该两数的下标（从0开始） .

func twoSum(input []int, target int) []int {
	var result []int
	for i := 0; i < len(input); i++ {
		for j := i + 1; j < len(input); j++ {
			if input[i]+input[j] == target {
				result = append(result, i)
				result = append(result, j)
				return result
			}
		}
	}
	return result
}

func main() {
	input := []int{2, 7, 8, 4, 9}
	target := 9
	fmt.Println(twoSum(input, target))
}
