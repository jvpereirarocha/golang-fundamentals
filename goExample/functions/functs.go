package main

import "fmt"

// func (params...) return_value

// The function below takes two ints and return their
// sum as an int
// Go requires explicit returns, it won't automatically
// return the value of the last expression
func plus(a int, b int) int {
	return a + b
}

// We can omit the type when they are the same
// In the function below all parameters are of the type int
func plusPlus(a, b, c int) int {
	return a + b + c
}

// The main function of the file
func main() {
	response := plus(1, 2)
	fmt.Println("1+2 =", response)

	response = plusPlus(1, 2, 3)
	fmt.Println("1+2+3 =", response)
}
