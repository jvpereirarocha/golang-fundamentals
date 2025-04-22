package main

import "fmt"

func main() {
	var a = "initial"
	fmt.Println(a)

	var b, c int = 1, 2
	fmt.Println(b, c)

	var d = true
	fmt.Println(d)

	var e int
	fmt.Println(e)

	/* The syntax below is available only inside functions */
	f := "apple"
	fmt.Println(f)

	var t string = "allow"
	fmt.Println("T value => ", t)

	z := "zeta"
	fmt.Printf("The letter is %s\n", z)
}
