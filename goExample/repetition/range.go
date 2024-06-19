package main

import "fmt"

func main() {
	nums := []int{2, 3, 4}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum: ", sum)

	// range on arrays ans slices provides both the index
	// and value for each entry.
	for i, num := range nums {
		if num == 3 {
			fmt.Printf("index: %v - value: %v\n", i, num)
		}
	}

	fmt.Println()

	// range on map iterates over key/value pairs
	kvs := map[string]string{"a": "apple", "b": "banana"}
	for key, value := range kvs {
		fmt.Printf("%s -> %s\n", key, value)
	}

	fmt.Println()

	//range can also iterate over just the keys of a map
	for key := range kvs {
		fmt.Printf("Key: %s\n", key)
	}

	fmt.Println()

	// Or can iterate over just the values of a map
	for _, value := range kvs {
		fmt.Printf("Value: %s\n", value)
	}

	// range on strings iterates over Unicode code points.
	// The first value is the starting byte index of the rune
	// and the second the rune itself
	for index, c := range "golang" {
		fmt.Println(index, c)
	}
}
