package main

import "fmt"

func factorial(n int) int {
	fmt.Println("Current value of n:", n)
	if n == 0 {
		return 1
	}
	return n * (factorial(n - 1))
}

func main() {
	fmt.Println(factorial(7))
	
	var fib func(n int) int
	
	fib = func(n int) int {
		fmt.Println("fib n value:", n)
		if n < 2 {
			return n
		}
		return fib(n - 1) + fib(n - 2)
	}
	
	fmt.Println(fib(7))
}