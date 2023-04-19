package main

// imports
import "fmt"

func main() {
	log := fmt.Println
	log("Loops In Go")

	// Simple Loop
	for i := 1; i <= 10; i++ {
		log(i)
	}
	// Syntax: (init)?; (condition); (finale)?

	// Infinite Loop
	/* for {} */

	// For loop as while loop
	w := 1
	for w < 5 {
		log(w)
		w++
	}

	// looping array
	a := []string{
		"Hello", "String", "Enjoy", "Hmm",
	}
	for value, index := range a {
		log(index, value)
	}
	fmt.Println("\n")

	// for string
	for index, character := range "Hello" {
		log(index, character)
		log(index, string(character))
	}

}
