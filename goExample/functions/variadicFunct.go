package main

import "fmt"

// Here's a function that will take an arbitrary number
// of 'ints' as arguments
// Within the function, the type of nums is equivalent
// to []int. We can call len(nums), iterate over it with
// range, etc
func sums(nums ...int) {
	fmt.Println(nums, " ")
	total := 0
	for _, num := range nums {
		total += num
	}
	fmt.Printf("Total: %v\n", total)
}

func main() {
	sums(1, 2)
	sums(1, 2, 3)

	// we can apply a slice to argument of the function sums
	nums := []int{4, 5, 6, 7, 8}
	sums(nums...)
}
