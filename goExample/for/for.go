package main

import "fmt"

func main() {
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	for j := 0; j < 3; j++ {
		fmt.Println(j)
	}

	// alternative notation 'for i := range 3' on version 1.22
	for i := range [3]int{} {
		fmt.Println("range ", i)
	}

	for {
		fmt.Println("loop")
		break
	}

	// alternative notation 'for i := range 6' on version 1.22
	for n := range [6]int{} {
		if n%2 == 0 {
			continue
		}
		fmt.Println((n))
	}
}
