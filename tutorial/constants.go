package main

import (
	"fmt"
	"math"
)

const s string = "constant"

func main() {
	fmt.Println("s: ", s)
	const n = 500000000
	
	const d = 3e20 / n
	fmt.Println("d: ", d)
	
	fmt.Println("int 64 n => ", int64(n))
	
	fmt.Println("math Sin: ", math.Sin(n))
}