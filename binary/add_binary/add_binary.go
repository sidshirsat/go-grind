// https://leetcode.com/problems/add-binary/

package main

import (
	"fmt"
	"strings"
)

func addBinary(a, b string) string {
	var result []int
	var carry int
	ptr1 := len(a) - 1
	ptr2 := len(b) - 1

	for ptr1 >= 0 || ptr2 >= 0 {
		var digit1, digit2 int

		if ptr1 >= 0 {
			digit1 = int(a[ptr1] - '0')
		} else {
			digit1 = 0
		}
		if ptr2 >= 0 {
			digit2 = int(b[ptr2] - '0')
		} else {
			digit2 = 0
		}

		aggregation := digit1 + digit2 + carry

		carry = aggregation / 2
		currentSum := aggregation % 2

		result = append(result, currentSum)

		ptr1--
		ptr2--
	}

	if carry != 0 {
		result = append(result, carry)
	}

	reversedResult := make([]string, len(result))

	for k, v := range result {
		reversedResult[len(result)-k-1] = fmt.Sprint(v)
	}

	binarySum := strings.Join(reversedResult, "")
	return binarySum
}

func main() {
	fmt.Println(addBinary("1010", "1011"))
}
