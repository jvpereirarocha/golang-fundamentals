package main

import "fmt"

/*
SlicesIndex takes a slice of any comparable type and an element
of that type and returns the index of the first occurrence of v in s,
or -1 if not present. The comparable constraint means that we can
compare values of this type with the == and != operators.
*/
func SlicesIndex[S ~[]E, E comparable](s S, v E) int {
	for i := range s {
		if v == s[i] {
			return i
		}
	}
	return -1
}

type Number interface {
	~int | ~float32
}

func Sum[T Number](a, b T) T{
	return a + b
}

type List[T any] struct {
	head, tail *element[T]
}

type element[T any] struct {
	next *element[T]
	val T
}

func (lst *List[T]) Push(v T) {
	if lst.tail == nil {
		lst.head = &element[T]{val: v}
		lst.tail = lst.head
	} else {
		lst.tail.next = &element[T]{val: v}
		lst.tail = lst.tail.next		
	}
}

func (lst *List[T]) AllElements() []T {
	var elems []T
	for e := lst.head; e != nil; e = e.next {
		elems = append(elems, e.val)	
	}
	return elems
}

func main() {
	var s = []string{"foo", "bar", "zoo"}
	fmt.Println("index of zoo:", SlicesIndex(s, "zoo"))

	_ = SlicesIndex[[]string, string](s, "zoo")
	lst := List[int]{}
	lst.Push(10)
	lst.Push(13)
	lst.Push(23)
	fmt.Println("list:", lst.AllElements())

	var a float32 = 2.2
	var b int = 3
	// The function Sum can receive float32 or int
	// But Golang doesn't allow use a mix of types.
	// You must use only int or only float32 (in this case)
	result1 := Sum(a, float32(b))
	fmt.Println("The result is =>", result1)
	
}