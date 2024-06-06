package main

import (
	"fmt"
)

func twoSum(nums []int, target int) []int {

	newslice := []int{}

	for index1 := 0; index1 < len(nums); index1++ {
		for index2 := index1 + 1; index2 < len(nums); index2++ {
			if nums[index1]+nums[index2] == target {
				return []int{index1, index2}
			}
		}
	}
	return newslice
}

/*
func explain what is happen
nums[index] + nums[index] == target
return slice [index1,index2]
*/
func main() {
	slice := []int{-1, -10, 19, -1}
	target := 9
	fmt.Println(twoSum(slice, target))
}
