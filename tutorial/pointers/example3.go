// Program to assign memory address to pointer

package main
import "fmt"

func main() {
	var name = "John"
	var ptr *string

	// assign the memory address of name to the pointer
	ptr = &name
	fmt.Println("Value of pointer (ptr) is", ptr)
	fmt.Println("Address of the variable 'name'", &name)

	// to get the value pointed by ptr
	fmt.Println("Value of ptr (before change it):", *ptr) // John

	// let's try to change the name variable value
	name = "Mary"
	fmt.Println("Value of variable 'name':", name)
	fmt.Println("Value of ptr (the pointer):", *ptr)

}