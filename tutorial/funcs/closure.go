package main

import "fmt"

func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {
	nextInt := intSeq()
	
	fmt.Println("First call:", nextInt())
	fmt.Println("Second call:", nextInt())
	fmt.Println("Third call:", nextInt())
	
	newInts := intSeq()
	fmt.Println("newInts call:", newInts())	
}