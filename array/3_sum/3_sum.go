// https://leetcode.com/problems/3sum/

package main

import (
	"fmt"
	"sort"
)

func findSumOfThree(nums []int, target int) bool {
	sort.Sort(sort.IntSlice(nums))

	low, high, aggregation := 0, 0, 0
	for i := 0; i < len(nums)-2; i++ {
		low = i + 1
		high = len(nums) - 1
		for low < high {
			aggregation = nums[i] + nums[low] + nums[high]

			if aggregation == target {
				fmt.Println("array is: ", nums)
				fmt.Println("triplets are: ", nums[i], nums[low], nums[high])
				return true
			} else if aggregation < target {
				low++
			} else {
				high--
			}
		}
	}
	return false
}

func main() {
	fmt.Println(findSumOfThree([]int{3, 7, 1, 2, 8, 4, 5, -2}, 1))
}
