/*
https://leetcode.com/problems/two-sum/
*/

package main

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
