package main

import "fmt"

type rect struct {
	width, height int
}

func (r *rect) area() int {
	return r.width * r.height
}

func (r rect) perim() int {
	return 2*r.width + 2*r.height
}

func main() {
	r := rect{width: 10, height: 5}
	
	fmt.Println("area 1: ", r.area())
	fmt.Println("perim 1:", r.perim())
	
	rp := &r
	fmt.Println("area 2: ", rp.area())
	fmt.Println("perim 2:", rp.perim())
}