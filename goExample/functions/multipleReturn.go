package main

import "fmt"

// The (int, bool) in the function signature
// shows that the function returns an int and a boolean
func onlyPositiveValues(a, b int) (int, bool) {
	hasError := false
	if a < 0 || b < 0 {
		hasError = true
	}
	sum := a + b
	return sum, hasError
}

func main() {
	a := -1
	b := 5
	sumOfNums, hasError := onlyPositiveValues(a, b)
	fmt.Printf("Sum of %v and %v: %v\n", a, b, sumOfNums)
	fmt.Printf("Has error: %v\n", hasError)
}
