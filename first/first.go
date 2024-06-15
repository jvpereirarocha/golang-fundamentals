package main

import (
	"fmt"
	"log"

	"example.com/greetings"
)

func main() {
	// Get a greeting message and print it
	log.SetPrefix("greetings: ")
	log.SetFlags(0)
	// message, err := greetings.Hello("Teste")
	names := []string{"Teste", "Qualquer", "Outro"}

	// message, err := greetings.Hello("")
	messages, err := greetings.Hellos(names)

	// If an error was returned, print it to the console
	// and exit the program
	if err != nil {
		log.Fatal(err)
	}

	// If no error was returned, print the returned message
	// to the console
	// fmt.Println(message)
	fmt.Println(messages)
}
