package main

import "fmt"

// ListNode definition for a singly-linked list
type ListNode struct {
	Val  int
	Next *ListNode
}

// mergeTwoLists merges two sorted linked lists and returns the head of the new list
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	// Create a dummy node to simplify handling of the merged list
	dummy := &ListNode{}
	current := dummy

	// Compare nodes from both lists and merge them
	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			current.Next = list1
			list1 = list1.Next
		} else {
			current.Next = list2
			list2 = list2.Next
		}
		current = current.Next
	}

	// Attach the remaining nodes from the non-empty list
	if list1 != nil {
		current.Next = list1
	} else if list2 != nil {
		current.Next = list2
	}

	// Return the merged list (skipping the dummy node)
	return dummy.Next
}

// Helper function to create a linked list from a slice
func createLinkedList(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}
	head := &ListNode{Val: nums[0]}
	current := head
	for _, val := range nums[1:] {
		current.Next = &ListNode{Val: val}
		current = current.Next
	}
	return head
}

// Helper function to print a linked list
func printLinkedList(head *ListNode) {
	for head != nil {
		fmt.Print(head.Val, " -> ")
		head = head.Next
	}
	fmt.Println("nil")
}

/*
func main() {
	// Create test cases
	list1 := createLinkedList([]int{1, 2, 4})
	list2 := createLinkedList([]int{1, 3, 4})

	// Merge the two sorted lists
	mergedList := mergeTwoLists(list1, list2)

	// Print the result
	printLinkedList(mergedList)
}
*/
