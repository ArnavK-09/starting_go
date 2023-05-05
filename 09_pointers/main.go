package main

import "fmt"

func main() {
	fmt.Println("Pointers in GO\n\t& - Reference Operater\n\t* - Dereference operator\n")

	// demo variables
	var a string = "Hellow"
	var b string = "World"
	fmt.Println(a, &a, b, &b) // variable refrence hex value

	// changing value with pointer func
	changeValue(&a, "Hello") // pass variable pointer
	fmt.Println(a, &a)       // address same

	// lol
	fmt.Println(*&b)

	// another way to create pointer
	c := new(uint) // init pointer 
	fmt.Println("Address: ", c, "Value: ", *c)

}

// pointer function
func changeValue(variable *string, value string) {
	*variable = value // derefrence with '*'
}
