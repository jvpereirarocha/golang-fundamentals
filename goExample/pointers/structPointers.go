package main

import "fmt"

type Person struct {
	name string
	age  int
}

// A closure to create a function on Person
// struct level
func (p *Person) test(name string) *Person {
	p.name = name
	return p
}

func newPerson(name string) *Person {
	person := Person{name: name}
	person.age = 27
	return &person
}

func main() {
	fmt.Println("Creating person...")
	p := newPerson("Joao Victor")
	fmt.Printf("Person: %s - Age: %v\n", p.name, p.age)

	test := p.test("JV")
	fmt.Printf("Name: %s - Age: %v\n", test.name, test.age)
}
