// https://leetcode.com/problems/two-sum/

package main

import "fmt"

// twoSum solves the Two Sum problem using a hashmap for O(n) efficiency.
func twoSum(nums []int, target int) []int {
	// Map to store the value and its corresponding index
	hashMap := make(map[int]int)

	// Iterate through the array
	for i, num := range nums {
		complement := target - num
		if idx, found := hashMap[complement]; found {
			return []int{idx, i}
		}
		hashMap[num] = i
	}

	// Return empty slice if no solution is found (shouldn't happen as per the problem guarantee)
	return nil
}

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9
	// Solve using twoSum
	result := twoSum(nums, target)
	// Print the result
	fmt.Printf("Indices: %v\n", result) // Expected Output: [0, 1]
}
