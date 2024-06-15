package main

import "fmt"

func main() {

	// Define a array with the syntax
	// var name [size]type
	var a [5]int
	fmt.Println("emp: ", a)

	a[4] = 100
	fmt.Println("set:", a)
	fmt.Println("get:", a[4])
	fmt.Printf("len: %v\n", len(a))

	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("dcl:", b)

	b = [...]int{3, 4, 5, 6, 7}
	fmt.Printf("dcl: %v\n", b)

	b = [...]int{100, 3: 400, 500} // index 0 will be 100 and index 3 will be 400
	fmt.Println("idx: ", b)

	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)

	twoD = [2][3]int{
		{1, 2, 3},
		{1, 2, 3},
	}
	fmt.Println("2d: ", twoD)
}
