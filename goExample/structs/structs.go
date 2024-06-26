package main

import "fmt"

type Person struct {
	name string
	age  int
}

func newPerson(name string) *Person {
	p := Person{name: name}
	p.age = 27
	// Yoou can safely return a pointer to local
	// variable as a local variable will
	// survive the scope of the function
	return &p
}

func main() {
	// This syntax creates a new struct
	fmt.Println(Person{"Bob", 20})

	// You can name the fields when initializing
	// a struct
	fmt.Println(Person{name: "Alice", age: 30})

	//Omitted fields will be zero-valued
	fmt.Println(Person{name: "Fred"})

	// An & prefix yields a pointer to the struct
	fmt.Println(&Person{name: "Ann", age: 40})

	// It's idiomatic to encapsulate new struct
	// creation in constructor functions
	fmt.Println(newPerson("John"))

	// Access struct fields with a dot
	s := Person{name: "Sean", age: 50}
	fmt.Println(s.name)

	// We can also use dots with struct pointers.
	// The pointers are automatically dereferenced
	sp := &s
	fmt.Println(sp.age)

	// Structs are mutable
	sp.age = 54
	fmt.Println(sp.age)
	fmt.Printf("Age after change: %v\n", s.age)

	// If a struct type is only used for a single
	// value, we don't have to give it a name.
	// The value can have an anonymous struct type.
	// This technique is commonly used for
	// table-driven tests
	dog := struct {
		name   string
		isGood bool
	}{
		"Rex",
		true,
	}
	fmt.Println(dog)
}
