package main

func isValid(s string) bool {
	// Stack to keep track of open parentheses
	stack := []rune{}

	// Map to match closing parentheses with their corresponding opening ones
	matching := map[rune]rune{
		')': '(',
		'}': '{',
		']': '[',
	}
	for _, char := range s {
		// If it's a closing bracket
		if open, ok := matching[char]; ok {
			// Check if the stack is empty or the top of the stack doesn't match
			if len(stack) == 0 || stack[len(stack)-1] != open {
				return false
			}
			// Pop the top of the stack
			stack = stack[:len(stack)-1]
		} else {
			// It's an opening bracket, push it onto the stack
			stack = append(stack, char)
		}
	}

	// If the stack is empty, all brackets matched correctly
	return len(stack) == 0
}

/* func main() {
	// Test cases
	fmt.Println(isValid("()"))     // true
	fmt.Println(isValid("()[]{}")) // true
	fmt.Println(isValid("(]"))     // false
	fmt.Println(isValid("([)]"))   // false
	fmt.Println(isValid("{[]}"))   // true
}
*/
