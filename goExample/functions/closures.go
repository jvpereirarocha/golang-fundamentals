package main

import "fmt"

// function intSeq returns another function, which
// we define anonymously in the body of intSeq.
// The returned function closes over the variable i to form
// a closure
func intSeq() func() int {
	i := 0
	fmt.Printf("i: %v\n", i)
	return func() int {
		i++
		return i
	}
}

func main() {
	// We call intSeq, assignin the result (a function)
	// to nextInt. THis function value captures its own
	// i value, which will be updated each time we call
	// nextInt
	nextInt := intSeq()

	// See the effect of the closure by calling nextInt
	// a few times
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	// Confirm that the state is unique creating and testing
	// a new one
	newInts := intSeq()
	fmt.Println(newInts())
}
